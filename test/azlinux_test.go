	"github.com/moby/buildkit/client/llb"
		SystemdDir: struct {
			Units   string
			Targets string
		}{
			Units:   "/usr/lib/systemd",
			Targets: "/etc/systemd/system",
		},
		SystemdDir: struct {
			Units   string
			Targets string
		}{
			Units:   "/usr/lib/systemd",
			Targets: "/etc/systemd/system",
		},
	SystemdDir  struct {
		Units   string
		Targets string
	}
				Build: map[string]dalec.PackageConstraints{"curl": {}},
		const src2Patch3File = "patch3"
		src2Patch3Content := []byte(`
diff --git a/file3 b/file3
new file mode 100700
index 0000000..5260cb1
--- /dev/null
+++ b/file3
@@ -0,0 +1,3 @@
+#!/usr/bin/env bash
+
+echo "Added another new file"
`)
		src2Patch3Context := llb.Scratch().File(
			llb.Mkfile(src2Patch3File, 0o600, src2Patch3Content),
		)
		src2Patch3ContextName := "patch-context"
