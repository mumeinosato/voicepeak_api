package src

import (
	"os"
	"os/exec"
	"strconv"
)

type Task struct {
	Text    string `json:"text"`
	Emotion Emotion `json:"emotion"`
}

type Emotion struct {
	Hightension int `json:"hightension"` //ハイテンション
	Livid       int `json:"livid"`       //怒り
	Lamenting   int `json:"lamenting"`   //悲しみ
	Despising   int `json:"despising"`   //さげすみ
}

func Taks(t Task) ([]byte, error) {
	mode := os.Getenv("MODE")

	output_path := ""
	if mode != "dev" {
		voicepeak_path := os.Getenv("VOICEPEAK_PATH")
		voicepeak_script_path := voicepeak_path + "/voicepeak"
		output_path = voicepeak_path + "/output.wav"

		emotion := "hightension=" + strconv.Itoa(t.Emotion.Hightension) + ",livid=" + strconv.Itoa(t.Emotion.Livid) + ",lamenting=" + strconv.Itoa(t.Emotion.Lamenting) + ",despising=" + strconv.Itoa(t.Emotion.Despising)
		out, err := exec.Command(voicepeak_script_path, "-s", t.Text, "-n", "Koharu Rikka", "-o", output_path, "-e", emotion).Output()
		if err != nil {
			return nil, err
		}
		println(string(out))
	} else {
		output_path = "./temp/temp.wav"
	}

	wavBytes, err := os.ReadFile(output_path)
	if err != nil {
		return nil, err
	}

	if mode != "dev" {
		err = os.Remove(output_path)
		if err != nil {
			return nil, err
		}
	}

	return wavBytes, nil
}
