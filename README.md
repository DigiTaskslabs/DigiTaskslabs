## Project Integration Process


1. Create a project first
   
![image](https://github.com/user-attachments/assets/f8b5b68d-eb74-4e10-99df-7590a3395582)
   
2. Publish the task
   
![image](https://github.com/user-attachments/assets/44dde22e-c196-4998-be5e-eba327274f4f)
   
3. Look at My Project, where the task you just published is in an unvalidated state
   
![image](https://github.com/user-attachments/assets/48460662-736b-4ac8-bc45-b62965e0921c)

![image](https://github.com/user-attachments/assets/8ce7561f-d49a-4482-ace2-6c7730d09abb)
   
4. Edit the task you just created, and you can get information about the verification task
   
![image](https://github.com/user-attachments/assets/b4670c47-1623-4f00-983e-87c49a66087e)

![image](https://github.com/user-attachments/assets/0115560f-fc52-4626-8333-8e34ddb9996b)
   
5. Send an authentication request to verify the task
   
![image](https://github.com/user-attachments/assets/8717127a-b13f-45a3-a7a4-ca8f9bbb2377)


 
## Call the API to settle the task

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
