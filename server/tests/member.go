package tests

import (
	"choco/server/internals/adapters"
	"choco/server/internals/models"
	"testing"
)

func testMemberAdapterAdd(t *testing.T, memberAdapter adapters.MemberAdapter, obj *models.Member) *models.Member {
	err := memberAdapter.Add(obj)

	if err != nil {
		t.Errorf("Not expected an error to create a member: %s", err)
	}

	return obj
}

func testMemberAdapterValidId(t *testing.T, memberAdapter adapters.MemberAdapter, validId string) *models.Member {
	member, err := memberAdapter.ID(validId)

	if err != nil {
		t.Errorf("Not expected an error to get a member by id that should exists in the database.")
	}

	return member
}

func testMemberAdapterInvalidId(t *testing.T, memberAdapter adapters.MemberAdapter, invalidId string) {
	_, err := memberAdapter.ID(invalidId)

	if err == nil {
		t.Errorf("Expected an error to get a member by id that shouldn't exists in the database, but got nil: %s", err)
	}
}

func testMemberAdapterValidCommunityId(t *testing.T, memberAdapter adapters.MemberAdapter, communityId string) []models.Member {
	members, err := memberAdapter.CommunityID(communityId)

	if err != nil {
		t.Errorf("Not expected an error to get the members of a community by the community's id: %s", err)
	}

	return members
}

func testMemberAdapterInvalidCommunityId(t *testing.T, memberAdapter adapters.MemberAdapter, invalidCommunityId string) {
	_, err := memberAdapter.CommunityID(invalidCommunityId)

	if err == nil {
		t.Errorf("Expected an error to get the members of a community by the community's id: %s", err)
	}
}

func testMemberAdapterValidUserId(t *testing.T, memberAdapter adapters.MemberAdapter, validUserId string) []models.Member {
	members, err := memberAdapter.UserID(validUserId)

	if err != nil {
		t.Errorf("Not expected an error to get the members that have a specific id: %s", err)
	}

	return members
}

func testMemberAdapterInvalidUserId(t *testing.T, memberAdapter adapters.MemberAdapter, invalidUserId string) {
	_, err := memberAdapter.UserID(invalidUserId)

	if err == nil {
		t.Errorf("Expected an error to get the members that have a specific id: %s", err)
	}
}

func testMemberAdapterValidUserIdAndCommunityId(t *testing.T, memberAdapter adapters.MemberAdapter, userId, communityId string) *models.Member {
	member, err := memberAdapter.MemberInTheCommunity(communityId, userId)

	if err != nil {
		t.Errorf("Not expected an error to get a member in the community even when the both id are corrects: %s", err)
	}

	return member
}

func testMemberAdapter(t *testing.T, memberAdapter adapters.MemberAdapter) {
	user, userErr := models.NewUser("okaabe", "okaabe@okaabe", []byte("okaabe"))

	if userErr != nil {
		t.Errorf("Not expected an error to create the user: %s", userErr)
	}

	community, communityErr := models.NewCommunity("test", "test", user.ID, false, false)

	if communityErr != nil {
		t.Errorf("Not expected an error to create the community: %s", communityErr)
	}

	member, memberErr := models.NewMember(user.ID, community.ID)

	if memberErr != nil {
		t.Errorf("Not expected an error to create the member: %s", memberErr)
	}

	testMemberAdapterAdd(t, memberAdapter, member)

	testMemberAdapterValidId(t, memberAdapter, member.UserID)
	testMemberAdapterInvalidId(t, memberAdapter, "a.a.a.a.a.a")

	testMemberAdapterInvalidCommunityId(t, memberAdapter, "ad.d..d.d.d.d")
	testMemberAdapterValidCommunityId(t, memberAdapter, member.CommunityID)

	testMemberAdapterInvalidUserId(t, memberAdapter, "ddwdwdw.dwdwdwd")
	testMemberAdapterValidUserId(t, memberAdapter, member.UserID)

	testMemberAdapterValidUserIdAndCommunityId(t, memberAdapter, member.UserID, member.CommunityID)
}
