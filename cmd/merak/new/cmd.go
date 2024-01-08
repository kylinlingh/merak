package new

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var pkg, branch string

func init() {
	NewCmd.Flags().StringVar(&pkg, "pkg", "", "项目包名")
	NewCmd.Flags().StringVar(&branch, "branch", "master", "项目模板分支")

	NewCmd.MarkFlagRequired("pkg")
}

// NewCmd 项目初始化工具
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "创建 helloworld 项目",
	Long:  `默认包名为 helloworld`,
	Run: func(cmd *cobra.Command, args []string) {
		color.White(strings.TrimLeft(`
MERAK
https://github.com/kylinlingh/sniper.git
`, "\n"))

		fail := false
		if err := exec.Command("git", "--version").Run(); err != nil {
			color.Red("git is not found")
			fail = true
		}

		//if err := exec.Command("make", "--version").Run(); err != nil {
		//	color.Red("make is not found")
		//	fail = true
		//}

		if fail {
			os.Exit(110)
		}

		//run("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@latest")

		parts := strings.Split(pkg, "/")
		path := parts[len(parts)-1]
		run("git", "clone", "https://github.com/kylinlingh/helloworld-job.git",
			"--quiet", "--depth=1", "--branch="+branch, path)

		if err := os.Chdir(path); err != nil {
			panic(err)
		}

		if pkg == "helloworld" {
			return
		}

		color.Cyan("rename helloworld to " + pkg)
		replace("go.mod", "module helloworld", "module "+pkg, 1)

		var files []string
		err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
			if strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".yml") {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			panic(err)
		}

		for _, p := range files {
			replace(p, `"helloworld`, `"`+pkg, -1)
		}

		color.Cyan("project created successfully")
	},
}

func replace(path, old, new string, n int) {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	s := string(b)
	s = strings.Replace(s, old, new, n)

	if err := os.WriteFile(path, []byte(s), 0); err != nil {
		panic(err)
	}
}

func run(name string, args ...string) {
	color.Cyan(name + " " + strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
