package worker

import "github.com/nanu-c/axolotl/app/config"

func (Api *TextsecureAPI) GetAvatarImage(id string) string {
	url := config.AttachDir + "/" + id

	return url
}
