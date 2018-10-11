package dcrm

import (
	"fmt"
	//get args
	"flag"
	"time"
	"os"
	//p2p
	//"github.com/fusion/go-fusion/p2p/dcrm"
	"github.com/fusion/go-fusion/log"
	//test dcrm bitInt
	"math/big"
)

//call define
func call(msg interface{}) {
	fmt.Printf("\n\ncall: msg=%v\n", msg)
}

func StartTest() {
	//callback
	//RegisterCallback(call)

	verbosity   := flag.Int("verbosity", int(log.LvlInfo), "log verbosity (0-9)")
	//get args
	//port := flag.Int("port", 0, "peer port")
	//bootnode := flag.String("bootnode", "", "bootstrap enode")
	//nodekeyfile := flag.String("nodekeyfile", "", "peer keyfile")
	//staticnodesfile := flag.String("staticnodes", "", "peer staticnodesfile")
	//flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	//glogger.Vmodule(*vmodule)
	log.Root().SetHandler(glogger)
	//fmt.Printf("port:%v, bootnode:%v, nodekeyfile:%v, staticnodesfile:%v\n", *port, *bootnode, *nodekeyfile, *staticnodesfile)

	//init p2p
	//dcrm.P2pInit(*port, *bootnode, *nodekeyfile, *staticnodesfile)

	if P2PTEST == 1 {
		select {}
	}

	time.Sleep(time.Duration(10) * time.Second)

	for {
		//get selfID, peers
		//count, peers := GetEnodes()
		//fmt.Printf("\ncount: %v, peers: %v\n", count, peers)

		time.Sleep(time.Duration(1) * time.Second)
		count, peers := GetGroup()
		fmt.Printf("group: %v, peers: %v\n", count, peers)
		if count > 0 {
			p := "hhhhhhhhhh"
			Send2Node(peers, &p)
		}
		//group: 4, peers: enode://74ec982620b1a9929b19e1373e74347289d43b8f6cd96dd03af8b72799a75139d601338dbb48e0786a304b28f35325407a0535625e1f8ead6f9292aeda0b4fd5@10.192.32.92:1236dcrmmsgenode://c25d9eb7e5100fc533a6507b0a2a1e1df027caa861d0c3b6ea1e6ade0fd17f3e932d1198dac2d13d02fdd894601b403d2ebe1a040d90eb409b7b68aad8c02e90@10.192.32.92:1237dcrmmsgenode://4538612f7d2b63aeea0adc96d550d3fa346c9abcdbde27e623dfb7a5c2977fdee0116047a826c89f20ef1d4fc44c6bba077e0039aa05b26608f081128a2780e6@10.192.32.92:1234dcrmmsgenode://3b2fd28db22477b9d3c5d51b12c36dc8f6af1e5b287e0de7353033ac6f3bd3ee3ae3d10256cc9af9d7af169749dfa09feaad03b398ba6c4c2a7ee559c11f8b13@10.192.32.92:1235
		//dcrm.SendMsg("test.........................")
	}
	select {}

	time.Sleep(time.Duration(5) * time.Second)

	//send
	ms := msg()
	SendMsg(ms)

	select {}
}

func msg() string {
	ss := ""
	for i := 0; i < 20; i++ {
		a, _ := new(big.Int).SetString("149538132774657726528707636360336072285983621283716341866783381657518370346851019858260832456828530008033228638611352429959811432602530865173229170389257564378211923685950755531505238837058174086920335663658007299016833930579190142733302787192962349263819218016641353162586263128086590197536551686649470858150350618785456971418349930832104986121457695940172259894785398687533391786467867726631938596486511367901516568978112757142419630344178113667535121256560750392040658835275511942192701455632856375360094654110506609442227926625226950851473547011845863799213473989320072877761242218539582755795868000263697836998474090588662414706888454874809674266627030310050419787442102412880785593702679692045546544678324521807290617899168862623088851045574478150024227834019720456951861075082019330308080142876318264220503114622414715335720216012228777904424800272212487065178597668582265177598601386221891513790771444030710626346787584945068392174449499454227939574351840134966158644998994253071165243081522707017502892682948657806254474909058265538220056912463355107659552522125535221371768932967262577221596192288246964658379224806800363453895132533539966950472006685897686030983111842328749543010305869598146974505445468053138588172334973", 10)

		s := a.Bytes()
		ss += string(s[:])
		ss += "dcrmmsg"
	}
	return ss
}