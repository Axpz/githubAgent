package github

import (
	"encoding/json"
	"testing"
)

const data = `{
	"ref": "refs/heads/main",
	"before": "dd070f12abcfa6fae6cd90660ccbe83426ea3b7d",
	"after": "b100b2a0e28f37287b874bc7da026f90aa70e608",
	"repository": {
		"id": 888323086,
		"node_id": "R_kgDONPK8Dg",
		"name": "githubAgent",
		"full_name": "Axpz/githubAgent",
		"private": false,
		"owner": {
			"name": "Axpz",
			"email": "axpzhang@gmail.com",
			"login": "Axpz",
			"id": 68174603,
			"node_id": "MDQ6VXNlcjY4MTc0NjAz",
			"avatar_url": "https://avatars.githubusercontent.com/u/68174603?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/Axpz",
			"html_url": "https://github.com/Axpz",
			"followers_url": "https://api.github.com/users/Axpz/followers",
			"following_url": "https://api.github.com/users/Axpz/following{/other_user}",
			"gists_url": "https://api.github.com/users/Axpz/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/Axpz/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/Axpz/subscriptions",
			"organizations_url": "https://api.github.com/users/Axpz/orgs",
			"repos_url": "https://api.github.com/users/Axpz/repos",
			"events_url": "https://api.github.com/users/Axpz/events{/privacy}",
			"received_events_url": "https://api.github.com/users/Axpz/received_events",
			"type": "User",
			"user_view_type": "public",
			"site_admin": false
		},
		"html_url": "https://github.com/Axpz/githubAgent",
		"description": null,
		"fork": false,
		"url": "https://github.com/Axpz/githubAgent",
		"forks_url": "https://api.github.com/repos/Axpz/githubAgent/forks",
		"keys_url": "https://api.github.com/repos/Axpz/githubAgent/keys{/key_id}",
		"collaborators_url": "https://api.github.com/repos/Axpz/githubAgent/collaborators{/collaborator}",
		"teams_url": "https://api.github.com/repos/Axpz/githubAgent/teams",
		"hooks_url": "https://api.github.com/repos/Axpz/githubAgent/hooks",
		"issue_events_url": "https://api.github.com/repos/Axpz/githubAgent/issues/events{/number}",
		"events_url": "https://api.github.com/repos/Axpz/githubAgent/events",
		"assignees_url": "https://api.github.com/repos/Axpz/githubAgent/assignees{/user}",
		"branches_url": "https://api.github.com/repos/Axpz/githubAgent/branches{/branch}",
		"tags_url": "https://api.github.com/repos/Axpz/githubAgent/tags",
		"blobs_url": "https://api.github.com/repos/Axpz/githubAgent/git/blobs{/sha}",
		"git_tags_url": "https://api.github.com/repos/Axpz/githubAgent/git/tags{/sha}",
		"git_refs_url": "https://api.github.com/repos/Axpz/githubAgent/git/refs{/sha}",
		"trees_url": "https://api.github.com/repos/Axpz/githubAgent/git/trees{/sha}",
		"statuses_url": "https://api.github.com/repos/Axpz/githubAgent/statuses/{sha}",
		"languages_url": "https://api.github.com/repos/Axpz/githubAgent/languages",
		"stargazers_url": "https://api.github.com/repos/Axpz/githubAgent/stargazers",
		"contributors_url": "https://api.github.com/repos/Axpz/githubAgent/contributors",
		"subscribers_url": "https://api.github.com/repos/Axpz/githubAgent/subscribers",
		"subscription_url": "https://api.github.com/repos/Axpz/githubAgent/subscription",
		"commits_url": "https://api.github.com/repos/Axpz/githubAgent/commits{/sha}",
		"git_commits_url": "https://api.github.com/repos/Axpz/githubAgent/git/commits{/sha}",
		"comments_url": "https://api.github.com/repos/Axpz/githubAgent/comments{/number}",
		"issue_comment_url": "https://api.github.com/repos/Axpz/githubAgent/issues/comments{/number}",
		"contents_url": "https://api.github.com/repos/Axpz/githubAgent/contents/{+path}",
		"compare_url": "https://api.github.com/repos/Axpz/githubAgent/compare/{base}...{head}",
		"merges_url": "https://api.github.com/repos/Axpz/githubAgent/merges",
		"archive_url": "https://api.github.com/repos/Axpz/githubAgent/{archive_format}{/ref}",
		"downloads_url": "https://api.github.com/repos/Axpz/githubAgent/downloads",
		"issues_url": "https://api.github.com/repos/Axpz/githubAgent/issues{/number}",
		"pulls_url": "https://api.github.com/repos/Axpz/githubAgent/pulls{/number}",
		"milestones_url": "https://api.github.com/repos/Axpz/githubAgent/milestones{/number}",
		"notifications_url": "https://api.github.com/repos/Axpz/githubAgent/notifications{?since,all,participating}",
		"labels_url": "https://api.github.com/repos/Axpz/githubAgent/labels{/name}",
		"releases_url": "https://api.github.com/repos/Axpz/githubAgent/releases{/id}",
		"deployments_url": "https://api.github.com/repos/Axpz/githubAgent/deployments",
		"created_at": 1731570414,
		"updated_at": "2025-01-18T16:05:57Z",
		"pushed_at": 1737216978,
		"git_url": "git://github.com/Axpz/githubAgent.git",
		"ssh_url": "git@github.com:Axpz/githubAgent.git",
		"clone_url": "https://github.com/Axpz/githubAgent.git",
		"svn_url": "https://github.com/Axpz/githubAgent",
		"homepage": null,
		"size": 9206,
		"stargazers_count": 1,
		"watchers_count": 1,
		"language": "Jupyter Notebook",
		"has_issues": true,
		"has_projects": true,
		"has_downloads": true,
		"has_wiki": true,
		"has_pages": false,
		"has_discussions": false,
		"forks_count": 0,
		"mirror_url": null,
		"archived": false,
		"disabled": false,
		"open_issues_count": 1,
		"license": null,
		"allow_forking": true,
		"is_template": false,
		"web_commit_signoff_required": false,
		"topics": [],
		"visibility": "public",
		"forks": 0,
		"open_issues": 1,
		"watchers": 1,
		"default_branch": "main",
		"stargazers": 1,
		"master_branch": "main"
	},
	"pusher": {
		"name": "Axpz",
		"email": "axpzhang@gmail.com"
	},
	"sender": {
		"login": "Axpz",
		"id": 68174603,
		"node_id": "MDQ6VXNlcjY4MTc0NjAz",
		"avatar_url": "https://avatars.githubusercontent.com/u/68174603?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/Axpz",
		"html_url": "https://github.com/Axpz",
		"followers_url": "https://api.github.com/users/Axpz/followers",
		"following_url": "https://api.github.com/users/Axpz/following{/other_user}",
		"gists_url": "https://api.github.com/users/Axpz/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/Axpz/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/Axpz/subscriptions",
		"organizations_url": "https://api.github.com/users/Axpz/orgs",
		"repos_url": "https://api.github.com/users/Axpz/repos",
		"events_url": "https://api.github.com/users/Axpz/events{/privacy}",
		"received_events_url": "https://api.github.com/users/Axpz/received_events",
		"type": "User",
		"user_view_type": "public",
		"site_admin": false
	},
	"created": false,
	"deleted": false,
	"forced": true,
	"base_ref": null,
	"compare": "https://github.com/Axpz/githubAgent/compare/dd070f12abcf...b100b2a0e28f",
	"commits": [
		{
			"id": "b100b2a0e28f37287b874bc7da026f90aa70e608",
			"tree_id": "ef547727888d970d110f5e654e16363c5b6ef413",
			"distinct": true,
			"message": "add redeployment\n\nSigned-off-by: zx <axpzhang@gmail.com>",
			"timestamp": "2025-01-19T00:15:50+08:00",
			"url": "https://github.com/Axpz/githubAgent/commit/b100b2a0e28f37287b874bc7da026f90aa70e608",
			"author": {
			"name": "zx",
				"email": "axpzhang@gmail.com",
				"username": "Axpz"
			},
			"committer": {
				"name": "zx",
				"email": "axpzhang@gmail.com",
				"username": "Axpz"
			},
			"added": [
				"docs/redeploy.sh"
			],
			"removed": [],
			"modified": [
				"internal/server/handlers/github/webhook_handler.go"
			]
		}
	],
	"head_commit": {
		"id": "b100b2a0e28f37287b874bc7da026f90aa70e608",
		"tree_id": "ef547727888d970d110f5e654e16363c5b6ef413",
		"distinct": true,
		"message": "add redeployment\n\nSigned-off-by: zx <axpzhang@gmail.com>",
		"timestamp": "2025-01-19T00:15:50+08:00",
		"url": "https://github.com/Axpz/githubAgent/commit/b100b2a0e28f37287b874bc7da026f90aa70e608",
		"author": {
			"name": "zx",
			"email": "axpzhang@gmail.com",
			"username": "Axpz"
		},
		"committer": {
			"name": "zx",
			"email": "axpzhang@gmail.com",
			"username": "Axpz"
		},
		"added": [
			"docs/redeploy.sh"
		],
		"removed": [],
		"modified": [
			"internal/server/handlers/github/webhook_handler.go"
		]
	}
}`

func TestParsePushEvent(t *testing.T) {
	var event PushEvent
	err := json.Unmarshal([]byte(data), &event)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if event.Ref != "refs/heads/main" {
		t.Errorf("Expected ref to be 'refs/heads/main', got '%s'", event.Ref)
	}

	if event.Before != "dd070f12abcfa6fae6cd90660ccbe83426ea3b7d" {
		t.Errorf("Expected before to be 'dd070f12abcfa6fae6cd90660ccbe83426ea3b7d', got '%s'", event.Before)
	}

	if event.After != "b100b2a0e28f37287b874bc7da026f90aa70e608" {
		t.Errorf("Expected after to be 'b100b2a0e28f37287b874bc7da026f90aa70e608', got '%s'", event.After)
	}

	if len(event.Commits) != 1 {
		t.Fatalf("Expected 1 commit, got %d", len(event.Commits))
	}

	if event.Commits[0].ID != "b100b2a0e28f37287b874bc7da026f90aa70e608" {
		t.Errorf("Expected commit ID to be 'b100b2a0e28f37287b874bc7da026f90aa70e608', got '%s'", event.Commits[0].ID)
	}

	if event.Commits[0].Message != "add redeployment\n\nSigned-off-by: zx <axpzhang@gmail.com>" {
		t.Errorf("Expected commit message to be 'add redeployment\n\nSigned-off-by: zx <axpzhang@gmail.com>', got '%s'", event.Commits[0].Message)
	}
}
