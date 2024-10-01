## Project Integration Process

1. **Create a new project**
2. **Create a new task**
3. **Select a task you created and go to edit the task**
4. **Copy the token**
   
![image](https://github.com/user-attachments/assets/14e6505e-b054-4975-87bf-9c418b42af7c)

5. Call the API to settle the task

```javascript
  const fetch = async () => {
    try {
      const url = 'https://testearn.gametop.me/earn/task/check-ref'
      const params = { uid: 1725584873 } // User id

      // add Authorization token to header
      const headers = { 'Content-Type': 'multipart/form-data', Authorization: 'Az8WyMRHyZdwTDtZcnYN6A==' }
      const { data } = await axios.post(url, params, { headers })
      // do something
    } catch (e) {
      console.error(e)
    }
  }
```
