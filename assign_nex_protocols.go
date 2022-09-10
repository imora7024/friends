package main

import (
	friends_3ds "github.com/PretendoNetwork/friends-secure/3ds"
	"github.com/PretendoNetwork/friends-secure/globals"
	friends_wiiu "github.com/PretendoNetwork/friends-secure/wiiu"
	nexproto "github.com/PretendoNetwork/nex-protocols-go"
)

func assignNEXProtocols() {
	secureServer := nexproto.NewSecureProtocol(globals.NEXServer)
	accountManagementServer := nexproto.NewAccountManagementProtocol(globals.NEXServer)
	friendsWiiUServer := nexproto.NewFriendsWiiUProtocol(globals.NEXServer)
	friends3DSServer := nexproto.NewFriends3DSProtocol(globals.NEXServer)

	// Account Management protocol handles
	accountManagementServer.NintendoCreateAccount(nintendoCreateAccount)

	// Secure protocol handles
	secureServer.Register(register)
	secureServer.RegisterEx(registerEx)

	// Friends (WiiU) protocol handles
	friendsWiiUServer.UpdateAndGetAllInformation(friends_wiiu.UpdateAndGetAllInformation)
	friendsWiiUServer.AddFriendRequest(friends_wiiu.AddFriendRequest)
	friendsWiiUServer.AcceptFriendRequest(friends_wiiu.AcceptFriendRequest)
	friendsWiiUServer.MarkFriendRequestsAsReceived(friends_wiiu.MarkFriendRequestsAsReceived)
	friendsWiiUServer.UpdatePresence(friends_wiiu.UpdatePresence)
	friendsWiiUServer.UpdateComment(friends_wiiu.UpdateComment)
	friendsWiiUServer.UpdatePreference(friends_wiiu.UpdatePreference)
	friendsWiiUServer.GetBasicInfo(friends_wiiu.GetBasicInfo)
	friendsWiiUServer.DeletePersistentNotification(friends_wiiu.DeletePersistentNotification)
	friendsWiiUServer.CheckSettingStatus(friends_wiiu.CheckSettingStatus)
	friendsWiiUServer.GetRequestBlockSettings(friends_wiiu.GetRequestBlockSettings)

	// Friends (3DS) protocol handles
	friends3DSServer.UpdateProfile(friends_3ds.UpdateProfile)
	friends3DSServer.UpdateMii(friends_3ds.UpdateMii)
	friends3DSServer.UpdatePreference(friends_3ds.UpdatePreference)
	friends3DSServer.SyncFriend(friends_3ds.SyncFriend)
	friends3DSServer.UpdatePresence(friends_3ds.UpdatePresence)
	friends3DSServer.UpdateFavoriteGameKey(friends_3ds.UpdateFavoriteGameKey)
	friends3DSServer.UpdateComment(friends_3ds.UpdateComment)
}