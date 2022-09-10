package main

import (
	"fmt"
	"os"
	"time"

	"github.com/PretendoNetwork/friends-secure/database"
	"github.com/PretendoNetwork/friends-secure/globals"
	nex "github.com/PretendoNetwork/nex-go"
)

func nexServerStart() {
	globals.NEXServer = nex.NewServer()
	globals.NEXServer.SetFragmentSize(900)
	globals.NEXServer.SetPrudpVersion(0)
	globals.NEXServer.SetKerberosKeySize(16)
	globals.NEXServer.SetKerberosPassword(os.Getenv("KERBEROS_PASSWORD"))
	globals.NEXServer.SetPingTimeout(20) // Maybe too long?
	globals.NEXServer.SetAccessKey("ridfebb9")

	globals.NEXServer.On("Data", func(packet *nex.PacketV0) {
		request := packet.RMCRequest()

		fmt.Println("==Friends - Secure==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("====================")
	})

	globals.NEXServer.On("Kick", func(packet *nex.PacketV0) {
		pid := packet.Sender().PID()
		delete(globals.ConnectedUsers, pid)

		lastOnline := nex.NewDateTime(0)
		lastOnline.FromTimestamp(time.Now())

		database.UpdateUserLastOnlineTime(pid, lastOnline)
		sendUserWentOfflineWiiUNotifications(packet.Sender())

		fmt.Println("Leaving")
	})

	globals.NEXServer.On("Ping", func(packet *nex.PacketV0) {
		fmt.Print("Pinged. Is ACK: ")
		fmt.Println(packet.HasFlag(nex.FlagAck))
	})

	globals.NEXServer.On("Connect", connect)

	assignNEXProtocols()

	globals.NEXServer.Listen(":60001")
}