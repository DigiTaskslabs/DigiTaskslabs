---
icon: merge
---

# Project Integration Process

> _If you encounter problems with the connection, you can communicate through Issues_

# API Overview 
DigiTasks provides the following three core API interfaces to help developers manage detailed information about user tasks:
## Task Completion Notification Interface (check-ref): 
&nbsp;&nbsp;Used to notify the DigiTasks platform that a task has been completed by a user, triggering the platform to distribute rewards accordingly.
## Task Completion User List Interface (tgidlist): 
&nbsp;&nbsp;Used to retrieve the Telegram user ID list of users who have completed a task.
## Individual Task User Status Query Interface (info):
&nbsp;&nbsp;Used to check if a specific Telegram user has completed a task.
# Task Process Description
## Task Creation and Token Allocation: 
&nbsp;&nbsp;After a task is created and approved by an administrator, the system generates a unique task token for each task. This token will be used for authentication in subsequent API calls.
## Task Completion Notification Process: 
&nbsp;&nbsp;When a user completes a task on the platform (e.g., registration, successful invitation, etc.), the front-end or back-end should notify the DigiTasks platform by calling the check-ref interface, indicating that the specified user has completed the task. DigiTasks will distribute rewards based on this notification.
## Task Completion Query and User Management: 
&nbsp;&nbsp;Developers can use the tgidlist interface to obtain the list of users who have completed the task. They can also use the info interface to check the completion status of individual users for a given task.
## API Verification: 
The task publisher account needs to go to the API editing page in the task editor, click "send" for verification. Once verified, the task can be automatically published. path:(Your task edit page →Click Access API →Send)


### Api interface for task completion

This interface is used to notify us after the user completes the task. If your task is to invite registration, please notify us after the user successfully completes the task


url:https://api.digitasks.cc/earn/task/check-ref

Request parameter:

&nbsp;&nbsp;uid:"Telegram user id" //Here you need to grab the telegram user id
 
Request header :

&nbsp;&nbsp;Authorization:"task token"
 
type:POST

#### Responce:
| code     | message| solution|
| -------- | -------- |-------- |
| 30020    | parameter error|Check the 'uid' 'Authorization' parameter |
| 30008    | task not opened yet|Check whether the task is approved|
| 30022    | task failed|The number of remaining tasks is insufficient|
| 30024    | task is accomplished or gold shortage| Quest completed or not enough gold|

example:
```javascript
curl -X POST "https://api.digitasks.cc/earn/task/check-ref" -H "Authorization: <task_token>" -H "Content-Type: application/json" -d {"uid":<telegram_user_id>}
```
```javascript
  const fetch = async () => {
    try {
      const url = 'https://api.digitasks.cc/earn/task/check-ref'
      const params = { uid: 1725584873 } // Telegram id

      // add Authorization token to header
      const headers = { 'Content-Type': 'application/json', Authorization: 'lp+7fijStQuf/xDsZEfliUs+X9b5AZQFF+2MOBoIEOk=' }
      const { data } = await axios.post(url, params, { headers })
      // do something
    } catch (e) {
      console.error(e)
    }
  }
```
You can go to the [api documentation](https://apifox.com/apidoc/shared-7e2b39e5-13b3-4a3e-a65b-729357dee1c1?pwd=888888) for other language examples

#### There are also miniapp registered membership examples in the main.go file



### Get a list of Telegram IDs that have been completed for the current task
url:https://api.digitasks.cc/earn/task/tgidlist

Request parameter:

&nbsp;&nbsp;page:"page number"
 
&nbsp;&nbsp;pageSize:"Number of items per page"   
 
Request header :

&nbsp;&nbsp;Authorization:"task token"
 
type:GET

example:
```javascript
curl -X GET "https://api.digitasks.cc/earn/task/tgidlist?page=1&pageSize=30" -H "Authorization: <task_token>" 
```
```javascript
import axios from 'axios'

const fetch = async () => {
  try {
    const url = 'https://api.digitasks.cc/earn/task/tgidlist?page=1&pageSize=30'

    // add Authorization token to header
    const headers = {
      'Content-Type': 'application/json',
      Authorization: 'y7WdddkJH7ztsWx3xNHN0ks+X9b5AZQFF+2MOBoIEOk=', // Your API Access
    }
    const { data } = await axios.get(url, { headers })
    console.log(data)
    // do something
  } catch (e) {
    console.error(e)
  }
}
```

### Get whether or not to complete this task through the user's Telegram ID
url:https://api.digitasks.cc/earn/task/info

Request parameter:

&nbsp;&nbsp;tgid:"Telegram user id"
 
Request header :

&nbsp;&nbsp;Authorization:"task token"
 
type:GET

example:
```javascript
curl -X GET "https://api.digitasks.cc/earn/task/info?tgid=1725584873" -H "Authorization: <task_token>" 
```
```javascript
import axios from 'axios'

const fetch = async () => {
  try {
    const url = 'https://api.digitasks.cc/earn/task/info'

    const params = { tgid: 1725584873 } // Your telegram id

    // add Authorization token to header
    const headers = {
      'Content-Type': 'application/json',
      Authorization: 'y7WdddkJH7ztsWx3xNHN0ks+X9b5AZQFF+2MOBoIEOk=', // Your API Access
    }
    const { data } = await axios.get(url, { headers, params })
    console.log(data)
    // do something
  } catch (e) {
    console.error(e)
  }
}
```


