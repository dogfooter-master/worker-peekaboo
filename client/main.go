package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"worker-peekaboo/peekaboo/pkg/grpc/pb"
	"worker-peekaboo/peekaboo/pkg/service"
)

var fs = flag.NewFlagSet("worker-peekaboo", flag.ExitOnError)
var gRpcAddr = fs.String("grpc-addr", "", "gRPC server address")

var tokens []string
var services []string
var conn *grpc.ClientConn
var ids map[int][]string

func main() {
	fs.Parse(os.Args[1:])

	if len(*gRpcAddr) == 0 {
		*gRpcAddr = service.GetConfigServerGrpc()
	}
	fmt.Fprintf(os.Stderr, "%v\n", *gRpcAddr)

	var opts []grpc.DialOption
	var err error
	opts = append(opts, grpc.WithInsecure())
	conn, err = grpc.Dial(*gRpcAddr, opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to dial: %v\n", err)
		return
	}
	defer conn.Close()

	services = append(services, "Peekaboo")
	services = append(services, "RefreshWindow")
	services = append(services, "StartStreaming")
	services = append(services, "EndStreaming")
	services = append(services, "ChangeQuality")
	services = append(services, "ChangeFps")
	services = append(services, "ChangeProperties")
	services = append(services, "MouseDown")
	services = append(services, "MouseDown2")
	services = append(services, "MouseUp")
	services = append(services, "MouseUp2")
	services = append(services, "MouseMove")
	services = append(services, "MouseMove2")



	services = append(services, "DragTest")
	//services = append(services, "ReadAllPatient")

	showUsage()

	lastService := "Peekaboo"
	for {
		var caseService string
		fmt.Fprintf(os.Stdout, "> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		line = strings.TrimLeft(line, " ")
		if len(line) == 0 {
			caseService = lastService
		} else {
			tokens = strings.Split(line, " ")
			caseService = tokens[0]
			lastService = caseService
		}
		caseService = lastService
		fmt.Fprintf(os.Stdout, "\n")
		switch caseService {
		case "Peekaboo":
			Peekaboo()
		case "RepeatPeekaboo":
			RepeatPeekaboo()
		case "RefreshWindow":
			RefreshWindow()
		case "StartStreaming":
			StartStreaming()
		case "EndStreaming":
			EndStreaming()
		case "ChangeQuality":
			ChangeQuality()
		case "ChangeFps":
			ChangeFps()
		case "ChangeProperties":
			ChangeProperties()
		case "MouseDown":
			MouseDown()
		case "MouseDown2":
			MouseDown2()
		case "MouseUp":
			MouseUp()
		case "MouseUp2":
			MouseUp2()
		case "MouseMove":
			MouseMove()
		case "MouseMove2":
			MouseMove2()

		case "DragTest":
			DragTest()
		case "h":
			showUsage()
		}

		for _, e := range tokens {
			fmt.Fprintf(os.Stderr, "%v ", e)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
}
func DragTest() {
	if len(tokens) < 2 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle>\n", GetFunctionName())
		return
	}

	handle := tokens[1]

	tokens = []string{}
	tokens = append(tokens, "MouseDown2")
	tokens = append(tokens, handle)
	tokens = append(tokens, "645")
	tokens = append(tokens, "320")
	MouseDown2()

	for i := 321; i > 180; i-- {
		tokens = []string{}
		tokens = append(tokens, "MouseMove")
		tokens = append(tokens, handle)
		tokens = append(tokens, "645")
		tokens = append(tokens, strconv.Itoa(i))
		MouseMove2()
		time.Sleep(time.Duration(1*time.Millisecond))
	}
	time.Sleep(time.Duration(1000*time.Millisecond))

	tokens = []string{}
	tokens = append(tokens, "MouseUp2")
	tokens = append(tokens, handle)
	tokens = append(tokens, "645")
	tokens = append(tokens, "180")
	MouseUp2()

}
func MouseMove2() {
	if len(tokens) < 4 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle> <x> <y>\n", GetFunctionName())
		return
	}

	h, _ := strconv.Atoi(tokens[1])
	x, _ := strconv.Atoi(tokens[2])
	y, _ := strconv.Atoi(tokens[3])
	message := pb.MouseMove2Request{
		Handle:  int32(h),
		X: int32(x),
		Y: int32(y),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.MouseMove2(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.MouseMove2Request error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func MouseMove() {
	if len(tokens) < 4 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle> <x> <y>\n", GetFunctionName())
		return
	}

	h, _ := strconv.Atoi(tokens[1])
	x, _ := strconv.ParseFloat(tokens[2], 32)
	y, _ := strconv.ParseFloat(tokens[3], 32)
	message := pb.MouseMoveRequest{
		Handle:  int32(h),
		X: float32(x),
		Y: float32(y),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.MouseMove(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.MouseMoveRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func MouseUp2() {
	if len(tokens) < 4 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle> <x> <y>\n", GetFunctionName())
		return
	}

	h, _ := strconv.Atoi(tokens[1])
	x, _ := strconv.Atoi(tokens[2])
	y, _ := strconv.Atoi(tokens[3])
	message := pb.MouseUp2Request{
		Handle:  int32(h),
		X: int32(x),
		Y: int32(y),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.MouseUp2(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.MouseUp2Request error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func MouseUp() {
	if len(tokens) < 4 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle> <x> <y>\n", GetFunctionName())
		return
	}

	h, _ := strconv.Atoi(tokens[1])
	x, _ := strconv.ParseFloat(tokens[2], 32)
	y, _ := strconv.ParseFloat(tokens[3], 32)
	message := pb.MouseUpRequest{
		Handle:  int32(h),
		X: float32(x),
		Y: float32(y),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.MouseUp(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.MouseUpRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func MouseDown2() {
	if len(tokens) < 4 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle> <x> <y>\n", GetFunctionName())
		return
	}

	h, _ := strconv.Atoi(tokens[1])
	x, _ := strconv.Atoi(tokens[2])
	y, _ := strconv.Atoi(tokens[3])
	message := pb.MouseDown2Request{
		Handle:  int32(h),
		X: int32(x),
		Y: int32(y),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.MouseDown2(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.MouseDown2Request error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func MouseDown() {
	if len(tokens) < 4 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle> <x> <y>\n", GetFunctionName())
		return
	}

	h, _ := strconv.Atoi(tokens[1])
	x, _ := strconv.ParseFloat(tokens[2], 32)
	y, _ := strconv.ParseFloat(tokens[3], 32)
	message := pb.MouseDownRequest{
		Handle:  int32(h),
		X: float32(x),
		Y: float32(y),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.MouseDown(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.MouseDownRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func ChangeProperties() {
	if len(tokens) < 5 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle> <label> <fps> <quality>\n", GetFunctionName())
		return
	}

	h, _ := strconv.Atoi(tokens[1])
	n, _ := strconv.Atoi(tokens[3])
	m, _ := strconv.Atoi(tokens[4])
	message := pb.ChangePropertiesRequest{
		Handle:  int32(h),
		Label:   tokens[2],
		Fps:     int32(n),
		Quality: int32(m),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.ChangeProperties(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.ChangePropertiesRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func ChangeFps() {
	if len(tokens) < 2 {
		fmt.Fprintf(os.Stderr, "%s: <service> <keyword>\n", GetFunctionName())
		return
	}

	n, _ := strconv.Atoi(tokens[1])
	message := pb.ChangeFpsRequest{
		Fps: int32(n),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.ChangeFps(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.ChangeFpsRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func ChangeQuality() {
	if len(tokens) < 2 {
		fmt.Fprintf(os.Stderr, "%s: <service> <keyword>\n", GetFunctionName())
		return
	}

	n, _ := strconv.Atoi(tokens[1])
	message := pb.ChangeQualityRequest{
		Quality: int32(n),
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.ChangeQuality(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.ChangeQualityRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}
func EndStreaming() {
	if len(tokens) < 1 {
		fmt.Fprintf(os.Stderr, "%s: <service> <keyword>\n", GetFunctionName())
		return
	}

	message := pb.EndStreamingRequest{}
	if len(tokens) > 1 {
		handle, _ := strconv.Atoi(tokens[1])
		message.Handle = int32(handle)
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.EndStreaming(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.EndStreamingRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}

func StartStreaming() {
	if len(tokens) < 3 {
		fmt.Fprintf(os.Stderr, "%s: <service> <handle> <label> [<fps> <quality>]\n", GetFunctionName())
		return
	}

	h, _ := strconv.Atoi(tokens[1])
	message := pb.StartStreamingRequest{
		Handle:  int32(h),
		Label:   tokens[2],
	}
	if len(tokens) > 3 {
		n, _ := strconv.Atoi(tokens[3])
		message.Fps = int32(n)
	}
	if len(tokens) > 4 {
		n, _ := strconv.Atoi(tokens[4])
		message.Quality = int32(n)
	}
	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.StartStreaming(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.StartStreamingRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}

func RefreshWindow() {
	if len(tokens) < 1 {
		fmt.Fprintf(os.Stderr, "%s: <service> <keyword>\n", GetFunctionName())
		return
	}

	message := pb.RefreshWindowsRequest{}
	if len(tokens) > 1 {
		message.Keyword = tokens[1]
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.RefreshWindows(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.RefreshWindowsRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}

func RepeatPeekaboo() {
	for {
		tokens = []string{"Peekaboo", "find", "MOMO"}
		Peekaboo()
		//go Peekaboo()
		//time.Sleep(60 * time.Millisecond)
	}
}

func Peekaboo() {
	if len(tokens) < 3 {
		fmt.Fprintf(os.Stderr, "%s: <service> <category> <keyword>\n", GetFunctionName())
		return
	}

	message := pb.PikabuRequest{
		Category: tokens[1],
		Keyword:  tokens[2],
	}

	fmt.Fprintf(os.Stderr, "> %v Request: \n", GetFunctionName())
	if j, err2 := json.MarshalIndent(message, "", " "); err2 != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err2)
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", string(j))
	}
	fmt.Fprintf(os.Stdout, "\n")

	defer timeTrack(time.Now(), GetFunctionName())
	c := pb.NewPeekabooClient(conn)
	reply, err := c.Pikabu(
		context.Background(),
		&message,
	)
	fmt.Fprintf(os.Stderr, "< %v Response: \n", GetFunctionName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "pb.PikabuRequest error: %v", err)
		return
	} else {
		if j, err2 := json.MarshalIndent(reply, "", " "); err2 != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err2)
		} else {
			fmt.Fprintf(os.Stderr, "%v\n", string(j))
		}
	}
}

func GetFunctionName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	tokens := strings.Split(f.Name(), ".")

	return tokens[len(tokens)-1]
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Fprintf(os.Stderr, "%s took %v\n", name, elapsed.Seconds())
}

func showUsage() {
	fmt.Fprintf(os.Stdout, "\n")
	for i, e := range services {
		fmt.Fprintf(os.Stdout, "%3d) %s\n", i, e)
	}
	fmt.Fprintf(os.Stdout, "\n")
}
func RandomId(i int) string {
	if ids == nil {
		ids = make(map[int][]string)
	}
	c := rand.Intn(100)
	id := ""
	if len(ids[i]) == 0 || c < 80 {
		id = bson.NewObjectId().Hex()
		ids[i] = append(ids[i], id)
	} else {
		n := rand.Intn(len(ids[i]))
		id = ids[i][n]
	}

	return id
}
func RandomDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
func GeneratePid() string {
	n := rand.Intn(100)
	if n < 30 {
		return ""
	}

	randomPid := randomdata.Alphanumeric(7)
	return randomPid
}
func RandomGender() string {
	n := rand.Intn(100)
	if n < 50 {
		return "M"
	}
	return "F"
}
func RandomEthnicity() string {
	ethnicityList := []string{
		"Caucasian", "Latino", "Asian", "African",
	}
	n := rand.Intn(len(ethnicityList) - 1)
	return ethnicityList[n]
}
func RandomKoreanFirstName() string {
	firstNameList := []string{
		"민준", "서준", "예준", "주원", "도윤", "시우", "지후", "현우", "준우", "지훈", "도현", "건우", "우진", "민재", "현준",
		"선우", "서진", "연우", "정우", "유준", "승현", "준혁", "승우", "지환", "승민", "시윤", "지우", "민성", "유찬", "준영", "진우",
		"시후", "지원", "은우", "윤우", "수현", "동현", "재윤", "민규", "시현", "태윤", "재원", "민우", "재민", "은찬", "한결", "윤호",
		"민찬", "시원", "성민", "성현", "수호", "준호", "승준", "현서", "재현", "시온", "지성", "태민", "태현", "민혁", "예성", "민호",
		"하율", "지안", "성준", "우빈", "지율", "정민", "규민", "윤성", "지한", "민석", "지민", "이준", "준", "준수", "서우", "은호",
		"은성", "예찬", "이안", "윤재", "율", "하람", "태양", "준희", "준성", "하진", "현수", "도훈", "승원", "정현", "건", "지완",
		"민준", "서준", "예준", "주원", "도윤", "시우", "지후", "현우", "준우", "지훈", "도현", "건우", "우진", "민재", "현준",
		"선우", "서진", "연우", "정우", "유준", "승현", "준혁", "승우", "지환", "승민", "시윤", "지우", "민성", "유찬", "준영", "진우",
		"시후", "지원", "은우", "윤우", "수현", "동현", "재윤", "민규", "시현", "태윤", "재원", "민우", "재민", "은찬", "한결", "윤호",
		"민찬", "시원", "성민", "성현", "수호", "준호", "승준", "현서", "재현", "시온", "지성", "태민", "태현", "민혁", "예성", "민호",
		"하율", "지안", "성준", "우빈", "지율", "정민", "규민", "윤성", "지한", "민석", "지민", "이준", "준", "준수", "서우", "은호",
		"은성", "예찬", "이안", "윤재", "율", "하람", "태양", "준희", "준성", "하진", "현수", "도훈", "승원", "정현", "건", "지완",
		"민준", "서준", "예준", "주원", "도윤", "시우", "지후", "현우", "준우", "지훈", "도현", "건우", "우진", "민재", "현준",
		"선우", "서진", "연우", "정우", "유준", "승현", "준혁", "승우", "지환", "승민", "시윤", "지우", "민성", "유찬", "준영", "진우",
		"시후", "지원", "은우", "윤우", "수현", "동현", "재윤", "민규", "시현", "태윤", "재원", "민우", "재민", "은찬", "한결", "윤호",
		"민찬", "시원", "성민", "성현", "수호", "준호", "승준", "현서", "재현", "시온", "지성", "태민", "태현", "민혁", "예성", "민호",
		"하율", "지안", "성준", "우빈", "지율", "정민", "규민", "윤성", "지한", "민석", "지민", "이준", "준", "준수", "서우", "은호",
		"은성", "예찬", "이안", "윤재", "율", "하람", "태양", "준희", "준성", "하진", "현수", "도훈", "승원", "정현", "건", "지완",
		"강민", "승호", "율", "준", "윤", "건", "봄", "현", "솔", "산", "별", "찬", "민", "설", "진", "원", "결", "환", "강", "은", "린", "훈",
		"겸", "혁", "단", "한", "슬", "빈", "선", "호", "수", "담", "유", "범", "연", "희", "신", "휘", "정", "온", "훤", "안", "비", "권", "영",
		"도", "완", "인", "운", "후", "헌", "솜", "랑", "성", "주", "우", "률", "경", "엘", "란", "레", "승", "송", "든", "리", "샘", "늘", "룩",
		"웅", "을", "용", "하", "들", "림", "본", "석", "빛", "욱", "명", "해", "상", "금", "람", "홍", "탄", "이", "아", "미", "나", "철", "륜",
		"국", "서", "규", "룬", "루", "곤", "름", "언", "지", "휼", "효",
	}
	n := rand.Intn(len(firstNameList) - 1)
	return firstNameList[n]
}
func RandomKoreanLastName() string {
	lastNameList := []string{
		"김", "이", "박", "최", "정", "강", "조", "윤", "장", "임", "한", "오", "서", "신", "권", "황", "안", "송", "전",
		"홍", "유", "고", "문", "양", "손", "배", "조", "백", "허", "유", "남", "심", "노", "정", "하", "곽", "성", "차",
		"주", "우", "구", "민", "유", "류", "나", "엄", "원", "천", "방", "공", "남궁", "황보", "모", "장", "기", "반",
		"명", "맹", "제", "탁", "국", "여", "어", "은", "구", "석", "사", "가", "시", "갈", "호", "설", "팽", "사공", "음",
	}
	n := rand.Intn(len(lastNameList) - 1)
	return lastNameList[n]
}
func RandomFullName() (string, string, string) {
	n := rand.Intn(100)
	if n < 50 {
		return RandomKoreanFirstName(), "", RandomKoreanLastName()
	}

	return randomdata.FirstName(randomdata.RandomGender), randomdata.FirstName(randomdata.RandomGender), randomdata.LastName()
}
func RandomPid() string {
	randomPid := randomdata.Alphanumeric(7)
	return randomPid
}
func RandomDiseaseType() string {
	diseaseTypeList := []string{
		"Acne", "Actinic_Keratosis", "Birthmarks", "Blisters", "Cherry_Angiomas", "Cold_Sores", "Dry_Skin", "Eczema",
		"Fungal_Nail_Infections", "Melasma", "Moles", "Psoriasis", "Rashes", "Rosacea", "Scabies", "Scars",
		"Shingles", "Skin_Allergies_-_Contact_Dermatitis", "Skin_Cancer_&_Treatments", "Summer_Skin_Irritants",
		"Vitiligo", "Warts",
	}
	n := rand.Intn(len(diseaseTypeList) - 1)
	return diseaseTypeList[n]
}
func RandomLocation() string {
	locationList := []string{
		"wrists", "forearms", "genitals", "legs", "face", "eyes", "mouth", "forehead", "nose", "ear", "groin", "breasts",
	}
	n := rand.Intn(len(locationList) - 1)
	return locationList[n]
}
func RandomTagList() []string {
	tagList := []string{
		"10", "20", "30", "40", "50", "60", "70", "남자", "여자", "심각함", "암",
	}
	c := rand.Intn(5)
	var t map[string]bool
	var l []string
	for i := 0; i < c; i++ {
		n := rand.Intn(len(tagList) - 1)
		if t[tagList[n]] == false {
			l = append(l, tagList[n])
		}
	}
	return l
}
