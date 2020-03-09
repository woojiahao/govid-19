package data

import (
  . "github.com/woojiahao/govid-19/pkg/utility"
  "gopkg.in/src-d/go-git.v4"
  "os"
)

const Path = "tmp"

// Load the repository into /tmp.
func Load() {
  _, err := git.PlainClone(Path, false, &git.CloneOptions{
    URL:      "https://github.com/CSSEGISandData/COVID-19.git",
    Progress: os.Stdout,
  })
  Check(err)
}

// Update the repository to the latest changes
func Update() {
  r, err := git.PlainOpen(Path)
  Check(err)

  w, err := r.Worktree()
  Check(err)

  err = w.Pull(&git.PullOptions{
    RemoteName: "origin",
    Progress:   os.Stdout,
  })
  Check(err)
}
