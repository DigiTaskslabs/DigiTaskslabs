---
icon: merge
---

# Project Integration Process

> _If you encounter problems with the connection, you can communicate through Issues_


### Call the API to settle the task
url:https://api.digitasks.cc/earn/task/check-ref

Request parameter:

&nbsp;&nbsp;uid:"Telegram user id"
 
Request header :

&nbsp;&nbsp;Authorization:"task token"
 
type:POST

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
 Example of ./start request API for Telegram bot

 Python

```Python
from telegram import InlineKeyboardButton, InlineKeyboardMarkup, Update
from telegram.ext import Updater, CommandHandler, CallbackQueryHandler, CallbackContext
import requests 

def start(update: Update, context: CallbackContext):
    #Let's say you click start to complete the task
    uid = update.message.from_user.id
    url = 'https://api.digitasks.cc/earn/task/check-ref'  
    form_data = {  
        'UID': uid
    } 
    headers = {  
        'Content-Type': 'application/json',   
        'Authorization': ' YOUR_TASK_TOKEN'   #Replace with your task token
    } 
    try:
        response = requests.post(url, data=form_data, headers=headers)  
        response.raise_for_status()
        if response.status_code == 200:
            response_data = response.json()  
            print(response_data)
    except requests.exceptions.RequestException as e:
        print(f"error：{e}")
    #post end
    keyboard = [
        [InlineKeyboardButton("Shop", callback_data='shop')],
        [InlineKeyboardButton("Game", callback_data='game')],
        [InlineKeyboardButton("Profile", callback_data='profile')],
    ]
    reply_markup = InlineKeyboardMarkup(keyboard)
    update.message.reply_text("Choose an option:", reply_markup=reply_markup)

def button(update: Update, context: CallbackContext):
    query = update.callback_query
    query.answer()

    if query.data == 'shop':
        # Load Shop Mini App
        query.message.reply_text("Loading Shop Mini App...")
    elif query.data == 'game':
        # Load Game Mini App
        query.message.reply_text("Loading Game Mini App...")
    elif query.data == 'profile':
        # Load Profile Mini App
        query.message.reply_text("Loading Profile Mini App...")

def main():
    updater = Updater("YOUR_API_TOKEN", use_context=True)
    dp = updater.dispatcher

    dp.add_handler(CommandHandler("start", start))
    dp.add_handler(CallbackQueryHandler(button))

    updater.start_polling()
    updater.idle()

if __name__ == '__main__':
    main()
```

 Node.js

```Javascript
const TelegramBot = require('node-telegram-bot-api');
const axios = require('axios');

const token = 'YOUR_API_TOKEN';
const bot = new TelegramBot(token, { polling: true });


bot.onText(/\/start/, async (msg) => {
    const uid = msg.from.id;

    const url = 'https://api.digitasks.cc/earn/task/check-ref';
    const form_data = { UID: uid };
    const headers = {
        'Content-Type': 'application/json',
        'Authorization': 'YOUR_TASK_TOKEN' // Replace with your task token
    };

    try {
        const response = await axios.post(url, form_data, { headers });
        if (response.status === 200) {
            console.log(response.data);
        }
    } catch (error) {
        console.error(`Error: ${error.message}`);
    }

    const keyboard = [
        [
            { text: "Shop", callback_data: 'shop' },
            { text: "Game", callback_data: 'game' },
            { text: "Profile", callback_data: 'profile' }
        ]
    ];

    const replyMarkup = {
        inline_keyboard: keyboard
    };

    bot.sendMessage(msg.chat.id, "Choose an option:", { reply_markup: replyMarkup });
});


bot.on('callback_query', (query) => {
    const { data, message } = query;

    if (data === 'shop') {
        bot.sendMessage(message.chat.id, "Loading Shop Mini App...");
    } else if (data === 'game') {
        bot.sendMessage(message.chat.id, "Loading Game Mini App...");
    } else if (data === 'profile') {
        bot.sendMessage(message.chat.id, "Loading Profile Mini App...");
    }
});

```
#### Responce:
| code     | message| solution|
| -------- | -------- |-------- |
| 30020    | parameter error|Check the 'uid' 'Authorization' parameter |
| 30008    | task not opened yet|Check whether the task is approved|
| 30022    | task failed|The number of remaining tasks is insufficient|
| 30024    | task is accomplished or gold shortage| Quest completed or not enough gold|


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


