# ⚙️ API access

### Call the API to settle the task

```javascript
const fetch = async () => {
  try {
    const url = 'https://api.digitasks.cc/earn/task/check-ref'
    const params = { uid: 1725584873 } // User id

    // add Authorization token to header
    const headers = { 'Content-Type': 'multipart/form-data', Authorization: 'lp+7fijStQuf/xDsZEfliUs+X9b5AZQFF+2MOBoIEOk=' }
    const { data } = await axios.post(url, params, { headers })
    // do something
  } catch (e) {
    console.error(e)
  }
}
```



### Get a list of Telegram IDs that have been completed for the current task

```javascript
import axios from 'axios'

const fetch = async () => {
  try {
    const url = 'https://api.digitasks.cc/earn/task/tgidlist?page=1&pageSize=30'

    // add Authorization token to header
    const headers = {
      'Content-Type': 'multipart/form-data',
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

```javascript
import axios from 'axios'

const fetch = async () => {
  try {
    const url = 'https://api.digitasks.cc/earn/task/info'

    const params = { tgid: 1725584873 } // Your telegram id

    // add Authorization token to header
    const headers = {
      'Content-Type': 'multipart/form-data',
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

## Example of ./start request API for Telegram bot

#### Python

```python
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
        'Content-Type': 'multipart/form-data',   
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

#### Node.js

```python
const TelegramBot = require('node-telegram-bot-api');
const axios = require('axios');

const token = 'YOUR_API_TOKEN';
const bot = new TelegramBot(token, { polling: true });


bot.onText(/\/start/, async (msg) => {
    const uid = msg.from.id;

    const url = 'https://api.digitasks.cc/earn/task/check-ref';
    const form_data = { UID: uid };
    const headers = {
        'Content-Type': 'multipart/form-data',
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

