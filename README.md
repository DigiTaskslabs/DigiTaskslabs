---
icon: merge
---

# Project Integration Process

> _If you encounter problems with the connection, you can communicate through Issues_

1.After the administrator approves the task, the task token is copied

2.Write api request code in your project

3.After the task is successfully launched, go to the task editing page and click Verify

4.Task on-line

### Api interface for task completion

This interface is used to notify us after the user completes the task. If your task is to invite registration, please notify us after the user successfully completes the task


url:https://api.digitasks.cc/earn/task/check-ref

Request parameter:

&nbsp;&nbsp;uid:"Telegram user id"
 
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


