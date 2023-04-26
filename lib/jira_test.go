package lib

import (
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_GetIssueById(t *testing.T) {
	// Need to set up jira mock
	//return

	err := godotenv.Load("../.env")
	if err != nil {
		t.Fail()
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{
  "expand": "renderedFields,names,schema,operations,editmeta,changelog,versionedRepresentations,customfield_10232.properties,customfield_10226.requestTypePractice",
  "id": "18506",
  "self": "https://1centre.atlassian.net/rest/api/2/issue/18506",
  "key": "BT-7021",
  "fields": {
    "statuscategorychangedate": "2021-11-10T14:39:25.161+1300",
    "parent": {
      "id": "18493",
      "key": "BT-7008",
      "self": "https://1centre.atlassian.net/rest/api/2/issue/18493",
      "fields": {
        "summary": "Email fixes - Batch 3",
        "status": {
          "self": "https://1centre.atlassian.net/rest/api/2/status/10501",
          "description": "This status is managed internally by Jira Software",
          "iconUrl": "https://1centre.atlassian.net/",
          "name": "Released",
          "id": "10501",
          "statusCategory": {
            "self": "https://1centre.atlassian.net/rest/api/2/statuscategory/3",
            "id": 3,
            "key": "done",
            "colorName": "green",
            "name": "Done"
          }
        },
        "priority": {
          "self": "https://1centre.atlassian.net/rest/api/2/priority/3",
          "iconUrl": "https://1centre.atlassian.net/images/icons/priorities/medium.svg",
          "name": "Medium",
          "id": "3"
        },
        "issuetype": {
          "self": "https://1centre.atlassian.net/rest/api/2/issuetype/10105",
          "id": "10105",
          "description": "Created by Jira Agile - do not edit or delete. Issue type for a user story.",
          "iconUrl": "https://1centre.atlassian.net/images/icons/issuetypes/story.svg",
          "name": "Story",
          "subtask": false,
          "hierarchyLevel": 0
        }
      }
    },
    "customfield_10230": null,
    "fixVersions": [
      {
        "self": "https://1centre.atlassian.net/rest/api/2/version/10208",
        "id": "10208",
        "description": "Included releases:\n4.44.7\n4.44.8\n4.44.9",
        "name": "4.44.6",
        "archived": false,
        "released": true,
        "releaseDate": "2021-11-16"
      }
    ],
    "customfield_10231": null,
    "customfield_10232": null,
    "resolution": {
      "self": "https://1centre.atlassian.net/rest/api/2/resolution/10000",
      "id": "10000",
      "description": "Work has been completed on this issue.",
      "name": "Done"
    },
    "customfield_10233": null,
    "customfield_10113": null,
    "customfield_10234": null,
    "customfield_10235": null,
    "customfield_10114": "1|hzz7qf:",
    "customfield_10225": null,
    "customfield_10226": null,
    "customfield_10227": [],
    "customfield_10228": null,
    "customfield_10229": null,
    "lastViewed": "2023-04-26T09:53:36.550+1200",
    "customfield_10220": null,
    "customfield_10221": null,
    "customfield_10100": "2021-10-31T15:22:41.084+1300",
    "priority": {
      "self": "https://1centre.atlassian.net/rest/api/2/priority/3",
      "iconUrl": "https://1centre.atlassian.net/images/icons/priorities/medium.svg",
      "name": "Medium",
      "id": "3"
    },
    "customfield_10101": "3_*:*_2_*:*_1733055632_*|*_10509_*:*_3_*:*_713784737_*|*_10200_*:*_1_*:*_785561_*|*_10100_*:*_2_*:*_4829138030_*|*_10501_*:*_3_*:*_479026_*|*_10500_*:*_2_*:*_424385849_*|*_10511_*:*_2_*:*_68827578",
    "customfield_10223": null,
    "customfield_10102": null,
    "labels": [],
    "customfield_10224": null,
    "aggregatetimeoriginalestimate": null,
    "timeestimate": null,
    "customfield_10219": null,
    "versions": [],
    "issuelinks": [],
    "assignee": {
      "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=602328456180010069ac44aa",
      "accountId": "602328456180010069ac44aa",
      "avatarUrls": {
        "48x48": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
        "24x24": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
        "16x16": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
        "32x32": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png"
      },
      "displayName": "James Besant",
      "active": false,
      "timeZone": "Pacific/Auckland",
      "accountType": "atlassian"
    },
    "status": {
      "self": "https://1centre.atlassian.net/rest/api/2/status/10501",
      "description": "This status is managed internally by Jira Software",
      "iconUrl": "https://1centre.atlassian.net/",
      "name": "Released",
      "id": "10501",
      "statusCategory": {
        "self": "https://1centre.atlassian.net/rest/api/2/statuscategory/3",
        "id": 3,
        "key": "done",
        "colorName": "green",
        "name": "Done"
      }
    },
    "components": [],
    "aggregatetimeestimate": null,
    "creator": {
      "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=5f6d623e4147d60077542e08",
      "accountId": "5f6d623e4147d60077542e08",
      "avatarUrls": {
        "48x48": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/48",
        "24x24": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/24",
        "16x16": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/16",
        "32x32": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/32"
      },
      "displayName": "Ben Bullock",
      "active": true,
      "timeZone": "Pacific/Auckland",
      "accountType": "atlassian"
    },
    "subtasks": [],
    "reporter": {
      "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=5f6d623e4147d60077542e08",
      "accountId": "5f6d623e4147d60077542e08",
      "avatarUrls": {
        "48x48": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/48",
        "24x24": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/24",
        "16x16": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/16",
        "32x32": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/32"
      },
      "displayName": "Ben Bullock",
      "active": true,
      "timeZone": "Pacific/Auckland",
      "accountType": "atlassian"
    },
    "aggregateprogress": {
      "progress": 0,
      "total": 0
    },
    "progress": {
      "progress": 0,
      "total": 0
    },
    "votes": {
      "self": "https://1centre.atlassian.net/rest/api/2/issue/BT-7021/votes",
      "votes": 0,
      "hasVoted": false
    },
    "worklog": {
      "startAt": 0,
      "maxResults": 20,
      "total": 0,
      "worklogs": []
    },
    "issuetype": {
      "self": "https://1centre.atlassian.net/rest/api/2/issuetype/10101",
      "id": "10101",
      "description": "A small piece of work that's part of a larger task.",
      "iconUrl": "https://1centre.atlassian.net/rest/api/2/universal_avatar/view/type/issuetype/avatar/10316?size=medium",
      "name": "Sub-task",
      "subtask": true,
      "avatarId": 10316,
      "hierarchyLevel": -1
    },
    "timespent": null,
    "project": {
      "self": "https://1centre.atlassian.net/rest/api/2/project/10100",
      "id": "10100",
      "key": "BT",
      "name": "Development",
      "projectTypeKey": "software",
      "simplified": false,
      "avatarUrls": {
        "48x48": "https://1centre.atlassian.net/rest/api/2/universal_avatar/view/type/project/avatar/10608",
        "24x24": "https://1centre.atlassian.net/rest/api/2/universal_avatar/view/type/project/avatar/10608?size=small",
        "16x16": "https://1centre.atlassian.net/rest/api/2/universal_avatar/view/type/project/avatar/10608?size=xsmall",
        "32x32": "https://1centre.atlassian.net/rest/api/2/universal_avatar/view/type/project/avatar/10608?size=medium"
      }
    },
    "aggregatetimespent": null,
    "resolutiondate": "2021-11-10T14:39:25.157+1300",
    "workratio": -1,
    "issuerestriction": {
      "issuerestrictions": {},
      "shouldDisplay": false
    },
    "watches": {
      "self": "https://1centre.atlassian.net/rest/api/2/issue/BT-7021/watchers",
      "watchCount": 3,
      "isWatching": false
    },
    "created": "2021-08-12T15:11:48.767+1200",
    "customfield_10260": null,
    "customfield_10261": null,
    "customfield_10262": null,
    "customfield_10263": null,
    "customfield_10264": null,
    "customfield_10265": null,
    "customfield_10266": null,
    "customfield_10267": null,
    "customfield_10258": null,
    "customfield_10259": null,
    "updated": "2021-11-16T10:08:45.071+1300",
    "timeoriginalestimate": null,
    "customfield_10250": null,
    "description": "Template Name: OWN - BC Reassigned\n\nEmail method: send_reassign_message\n\nTemplate_id: TBC\n\nEmail log name: N/A\n\nDescription: Email to the *reassigned user (owner)* when an application is *reassigned* by the supplier.\n\nTASK: Create a new email\n\n\n\nScenario\n\nGiven: Application *is not* archived.\n\nOr: Application *is not* final approved.\n\nOr: Application *is not* final declined.\n\nWhen: Application is reassigned by the supplier\n\nThen: send email “OWN - BC Reassigned” to the *reassigned user (owner).*\n\n\n\nSUPPRESSION = possible",
    "customfield_10251": null,
    "customfield_10252": null,
    "customfield_10253": null,
    "customfield_10255": null,
    "customfield_10256": null,
    "timetracking": {},
    "customfield_10257": null,
    "customfield_10006": null,
    "customfield_10248": null,
    "security": null,
    "customfield_10007": {
      "hasEpicLinkFieldDependency": false,
      "showField": false,
      "nonEditableReason": {
        "reason": "PLUGIN_LICENSE_ERROR",
        "message": "The Parent Link is only available to Jira Premium users."
      }
    },
    "customfield_10249": null,
    "attachment": [],
    "summary": "Sendgrid template: OWN - BC Reassigned",
    "customfield_10240": null,
    "customfield_10241": null,
    "customfield_10242": null,
    "customfield_10000": "{}",
    "customfield_10001": null,
    "customfield_10243": null,
    "customfield_10244": null,
    "customfield_10002": null,
    "customfield_10245": null,
    "customfield_10236": null,
    "customfield_10237": null,
    "customfield_10117": null,
    "customfield_10238": null,
    "environment": null,
    "customfield_10239": null,
    "duedate": null,
    "comment": {
      "comments": [
        {
          "self": "https://1centre.atlassian.net/rest/api/2/issue/18506/comment/49097",
          "id": "49097",
          "author": {
            "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=557058%3Aa86d9106-2969-4738-83eb-186c5dab4cd5",
            "accountId": "557058:a86d9106-2969-4738-83eb-186c5dab4cd5",
            "avatarUrls": {
              "48x48": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/48",
              "24x24": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/24",
              "16x16": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/16",
              "32x32": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/32"
            },
            "displayName": "Miriana Lowrie",
            "active": true,
            "timeZone": "Pacific/Auckland",
            "accountType": "atlassian"
          },
          "body": "#1 Feedback\n\n* SME account\n* Reassigned to another user\n* Just keeps spinning and not allocating\n\n[https://uat.1centre.com/dashboard/applications/12223900-197a-4729-a132-e20c1b075efa/reallocate_reassign|https://uat.1centre.com/dashboard/applications/12223900-197a-4729-a132-e20c1b075efa/reallocate_reassign]",
          "updateAuthor": {
            "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=557058%3Aa86d9106-2969-4738-83eb-186c5dab4cd5",
            "accountId": "557058:a86d9106-2969-4738-83eb-186c5dab4cd5",
            "avatarUrls": {
              "48x48": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/48",
              "24x24": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/24",
              "16x16": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/16",
              "32x32": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/32"
            },
            "displayName": "Miriana Lowrie",
            "active": true,
            "timeZone": "Pacific/Auckland",
            "accountType": "atlassian"
          },
          "created": "2021-10-31T15:22:41.084+1300",
          "updated": "2021-10-31T15:22:41.084+1300",
          "jsdPublic": true
        },
        {
          "self": "https://1centre.atlassian.net/rest/api/2/issue/18506/comment/49098",
          "id": "49098",
          "author": {
            "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=557058%3Aa86d9106-2969-4738-83eb-186c5dab4cd5",
            "accountId": "557058:a86d9106-2969-4738-83eb-186c5dab4cd5",
            "avatarUrls": {
              "48x48": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/48",
              "24x24": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/24",
              "16x16": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/16",
              "32x32": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/32"
            },
            "displayName": "Miriana Lowrie",
            "active": true,
            "timeZone": "Pacific/Auckland",
            "accountType": "atlassian"
          },
          "body": "#2 Feedback\n\n* Enterprise account\n* Reassigned from website to another user\n* Just kept spinning so I went out of the VCF and back into the sales pipeline\n* I then went back into the VCF and it had transferred to the new user\n* The email triggered\n* But the email link is not working\n\n[https://uat.1centre.com/dashboard/applications/00ff9fcc-abe3-4521-bb66-68203d130b8f/reallocate_reassign|https://uat.1centre.com/dashboard/applications/00ff9fcc-abe3-4521-bb66-68203d130b8f/reallocate_reassign]\n\n[~accountid:602328456180010069ac44aa] ",
          "updateAuthor": {
            "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=557058%3Aa86d9106-2969-4738-83eb-186c5dab4cd5",
            "accountId": "557058:a86d9106-2969-4738-83eb-186c5dab4cd5",
            "avatarUrls": {
              "48x48": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/48",
              "24x24": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/24",
              "16x16": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/16",
              "32x32": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/557058:a86d9106-2969-4738-83eb-186c5dab4cd5/61bc8594-ccdc-4989-be24-0586920e8d2e/32"
            },
            "displayName": "Miriana Lowrie",
            "active": true,
            "timeZone": "Pacific/Auckland",
            "accountType": "atlassian"
          },
          "created": "2021-10-31T15:39:50.572+1300",
          "updated": "2021-10-31T15:39:50.572+1300",
          "jsdPublic": true
        },
        {
          "self": "https://1centre.atlassian.net/rest/api/2/issue/18506/comment/49143",
          "id": "49143",
          "author": {
            "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=602328456180010069ac44aa",
            "accountId": "602328456180010069ac44aa",
            "avatarUrls": {
              "48x48": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
              "24x24": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
              "16x16": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
              "32x32": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png"
            },
            "displayName": "James Besant",
            "active": false,
            "timeZone": "Pacific/Auckland",
            "accountType": "atlassian"
          },
          "body": "[~accountid:557058:a86d9106-2969-4738-83eb-186c5dab4cd5] [~accountid:5f6d623e4147d60077542e08] Have fixed the application link and deployed to staging for this one. Tests are looking right",
          "updateAuthor": {
            "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=602328456180010069ac44aa",
            "accountId": "602328456180010069ac44aa",
            "avatarUrls": {
              "48x48": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
              "24x24": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
              "16x16": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png",
              "32x32": "https://secure.gravatar.com/avatar/49165ffd1a367de1b3abc663b679ae7e?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FJB-2.png"
            },
            "displayName": "James Besant",
            "active": false,
            "timeZone": "Pacific/Auckland",
            "accountType": "atlassian"
          },
          "created": "2021-11-01T14:32:33.878+1300",
          "updated": "2021-11-01T14:32:33.878+1300",
          "jsdPublic": true
        },
        {
          "self": "https://1centre.atlassian.net/rest/api/2/issue/18506/comment/49177",
          "id": "49177",
          "author": {
            "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=5f6d623e4147d60077542e08",
            "accountId": "5f6d623e4147d60077542e08",
            "avatarUrls": {
              "48x48": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/48",
              "24x24": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/24",
              "16x16": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/16",
              "32x32": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/32"
            },
            "displayName": "Ben Bullock",
            "active": true,
            "timeZone": "Pacific/Auckland",
            "accountType": "atlassian"
          },
          "body": "looking good",
          "updateAuthor": {
            "self": "https://1centre.atlassian.net/rest/api/2/user?accountId=5f6d623e4147d60077542e08",
            "accountId": "5f6d623e4147d60077542e08",
            "avatarUrls": {
              "48x48": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/48",
              "24x24": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/24",
              "16x16": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/16",
              "32x32": "https://avatar-management--avatars.us-west-2.prod.public.atl-paas.net/5f6d623e4147d60077542e08/126877b3-8f7d-4223-95d1-3eaed577654d/32"
            },
            "displayName": "Ben Bullock",
            "active": true,
            "timeZone": "Pacific/Auckland",
            "accountType": "atlassian"
          },
          "created": "2021-11-02T08:13:14.365+1300",
          "updated": "2021-11-02T08:13:14.365+1300",
          "jsdPublic": true
        }
      ],
      "self": "https://1centre.atlassian.net/rest/api/2/issue/18506/comment",
      "maxResults": 4,
      "total": 4,
      "startAt": 0
    }
  }
}`)
	}))
	defer ts.Close()

	jiraConfig := JiraConfig{
		Username:      os.Getenv("JIRA_USERNAME"),
		Password:      os.Getenv("JIRA_PASSWORD"),
		Url:           ts.URL,
		ProjectPrefix: os.Getenv("JIRA_PROJECT_PREFIX"),
	}

	issue := jiraConfig.GetIssueById("BT-7021")
	if issue.Err != nil {
		t.Fail()
	}
	println(issue.Status)
	if issue.Summary == "" {
		t.Error("Summary is blank")
	}
}
