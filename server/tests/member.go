package tests

import (
	"choco/server/internals/adapters"
	"choco/server/internals/models"
	"testing"
)

func testMemberAdapterAdd(t *testing.T, memberAdapter adapters.MemberAdapter) *models.Member {
	return nil
}

func testMemberAdapterValidId(t *testing.T, memberAdapter adapters.MemberAdapter, validId string) *models.Member {
	return nil
}

func testMemberAdapterInvalidId(t *testing.T, memberAdapter adapters.MemberAdapter, invalidId string) {

}

func testMemberAdapterValidCommunityId(t *testing.T, memberAdapter adapters.MemberAdapter, communityId string) []models.Member {
	return nil
}

func testMemberAdapterInvalidCommunityId(t *testing.T, memberAdapter adapters.MemberAdapter, invalidCommunityId string) {

}

func testMemberAdapterValidUserId(t *testing.T, memberAdapter adapters.MemberAdapter, validUserId string) []models.Member {
	return nil
}

func testMemberAdapterInvalidUserId(t *testing.T, memberAdapter adapters.MemberAdapter, invalidUserId string) {

}

func testMemberAdapterValidUserIdAndCommunityId(t *testing.T, memberAdapter adapters.MemberAdapter, userId, communityId string) *models.User {

	return nil
}

func testMemberAdapter(t *testing.T, memberAdapter adapters.MemberAdapter) {
	testMemberAdapterAdd(t, memberAdapter)
}
