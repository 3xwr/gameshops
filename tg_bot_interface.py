import requests
import re
import os
from aiogram import Bot, types
from aiogram.dispatcher import Dispatcher
from aiogram.utils import executor
from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton
from config import token

bot = Bot(token=token)
dp = Dispatcher(bot)

@dp.message_handler(commands=['hello'])
async def process_start_command(message: types.Message):
    await bot.send_message(message.from_user.id, "Привет!\nУ нас тут есть поиск игр по разным торговым площадкам")


@dp.message_handler(commands=['start'])
async def get_start(message: types.Message):
    await message.answer("Введи название игры, которую хочешь приобрести", parse_mode="Markdown")


@dp.message_handler()
async def get_name(message: types.Message):
    r = requests.get(f'http://127.0.0.1:5555/getname?name={message.text}')
    if len(r.json()) == 1:
        r = requests.get('http://127.0.0.1:5555/compareprice?name='+message.text)
        for i in range(0, 3):
            if not (r.json()[0].get('status')):
                p = requests.get(r.json()[0]['store_image'])
                out = open(str(message.message_id) + ".jpg", "wb")
                out.write(p.content)
                out.close()
                i = 4
        third_kb = InlineKeyboardMarkup()
        if not(r.json()[0].get('status')):
            button1 = InlineKeyboardButton("Steam : "+r.json()[0]['store_price'], url = r.json()[0]['store_app_url'])
            third_kb.add(button1)
        if not (r.json()[1].get('status')):
            button2 = InlineKeyboardButton("SteamPay : "+r.json()[1]['store_price'], url = r.json()[1]['store_app_url'])
            third_kb.add(button2)
        if not (r.json()[2].get('status')):
            button3 = InlineKeyboardButton("GOG : "+r.json()[2]['store_price'], url = r.json()[2]['store_app_url'])
            third_kb.add(button3)
        if not (r.json()[3].get('status')):
            button4 = InlineKeyboardButton("Плати.ру : "+r.json()[3]['store_price'], url = r.json()[3]['store_app_url'])
            third_kb.add(button4)

        await bot.send_photo(message.from_user.id, photo=open(str(message.message_id)+".jpg", "rb"), reply_markup=third_kb)
        os.remove(str(message.message_id)+".jpg")
    else:
        btn1 = InlineKeyboardButton(r.json()[0],callback_data="first")
        btn2 = InlineKeyboardButton(r.json()[1], callback_data="second")
        btn3 = InlineKeyboardButton(r.json()[2], callback_data="third")
        second_kb = InlineKeyboardMarkup()
        second_kb.add(btn1)
        second_kb.add(btn2)
        second_kb.add(btn3)
        await bot.send_message(message.from_user.id, 'Уточни название: '+str(message.text), reply_markup=second_kb)

@dp.callback_query_handler(text='first')
async def first_price(callback_query: types.CallbackQuery):
    text = re.sub(r'Уточни название: ','',callback_query.message.text)
    j = requests.get(f'http://127.0.0.1:5555/getname?name={text}')
    r = requests.get('http://127.0.0.1:5555/compareprice?name=' + j.json()[0])
    for i in range(0,3):
        if not (r.json()[0].get('status')):
            p = requests.get(r.json()[0]['store_image'])
            out = open(str(callback_query.message.message_id) + ".jpg", "wb")
            out.write(p.content)
            out.close()
            i=4

    third_kb = InlineKeyboardMarkup()
    if not (r.json()[0].get('status')):
        button1 = InlineKeyboardButton("Steam : " + r.json()[0]['store_price'], url=r.json()[0]['store_app_url'])
        third_kb.add(button1)
    if not (r.json()[1].get('status')):
        button2 = InlineKeyboardButton("SteamPay : " + r.json()[1]['store_price'], url=r.json()[1]['store_app_url'])
        third_kb.add(button2)
    if not (r.json()[2].get('status')):
        button3 = InlineKeyboardButton("GOG : " + r.json()[2]['store_price'], url=r.json()[2]['store_app_url'])
        third_kb.add(button3)
    if not (r.json()[3].get('status')):
        button4 = InlineKeyboardButton("Плати.ру : " + r.json()[3]['store_price'], url=r.json()[3]['store_app_url'])
        third_kb.add(button4)
    await bot.delete_message(chat_id=callback_query.from_user.id, message_id=callback_query.message.message_id)
    await bot.send_photo(callback_query.from_user.id, photo=open(str(callback_query.message.message_id) + ".jpg", "rb"),reply_markup=third_kb)
    os.remove(str(callback_query.message.message_id) + ".jpg")

@dp.callback_query_handler(text='second')
async def second_price(callback_query: types.CallbackQuery):
    text = re.sub(r'Уточни название: ','',callback_query.message.text)
    j = requests.get(f'http://127.0.0.1:5555/getname?name={text}')
    r = requests.get('http://127.0.0.1:5555/compareprice?name=' + j.json()[1])
    for i in range(0,3):
        if not (r.json()[0].get('status')):
            p = requests.get(r.json()[0]['store_image'])
            out = open(str(callback_query.message.message_id) + ".jpg", "wb")
            out.write(p.content)
            out.close()
            i=4
    third_kb = InlineKeyboardMarkup()
    if not (r.json()[0].get('status')):
        button1 = InlineKeyboardButton("Steam : " + r.json()[0]['store_price'], url=r.json()[0]['store_app_url'])
        third_kb.add(button1)
    if not (r.json()[1].get('status')):
        button2 = InlineKeyboardButton("SteamPay : " + r.json()[1]['store_price'], url=r.json()[1]['store_app_url'])
        third_kb.add(button2)
    if not (r.json()[2].get('status')):
        button3 = InlineKeyboardButton("GOG : " + r.json()[2]['store_price'], url=r.json()[2]['store_app_url'])
        third_kb.add(button3)
    if not (r.json()[3].get('status')):
        button4 = InlineKeyboardButton("Плати.ру : " + r.json()[3]['store_price'], url=r.json()[3]['store_app_url'])
        third_kb.add(button4)
    await bot.delete_message(chat_id=callback_query.from_user.id, message_id=callback_query.message.message_id)
    await bot.send_photo(callback_query.from_user.id, photo=open(str(callback_query.message.message_id) + ".jpg", "rb"),reply_markup=third_kb)
    os.remove(str(callback_query.message.message_id) + ".jpg")

@dp.callback_query_handler(text='third')
async def third_price(callback_query: types.CallbackQuery):
    text = re.sub(r'Уточни название: ','',callback_query.message.text)
    j = requests.get(f'http://127.0.0.1:5555/getname?name={text}')
    r = requests.get('http://127.0.0.1:5555/compareprice?name=' + j.json()[2])
    for i in range(0,3):
        if not (r.json()[0].get('status')):
            p = requests.get(r.json()[0]['store_image'])
            out = open(str(callback_query.message.message_id) + ".jpg", "wb")
            out.write(p.content)
            out.close()
            i=4
    third_kb = InlineKeyboardMarkup()
    if not (r.json()[0].get('status')):
        button1 = InlineKeyboardButton("Steam : " + r.json()[0]['store_price'], url=r.json()[0]['store_app_url'])
        third_kb.add(button1)
    if not (r.json()[1].get('status')):
        button2 = InlineKeyboardButton("SteamPay : " + r.json()[1]['store_price'], url=r.json()[1]['store_app_url'])
        third_kb.add(button2)
    if not (r.json()[2].get('status')):
        button3 = InlineKeyboardButton("GOG : " + r.json()[2]['store_price'], url=r.json()[2]['store_app_url'])
        third_kb.add(button3)
    if not (r.json()[3].get('status')):
        button4 = InlineKeyboardButton("Плати.ру : " + r.json()[3]['store_price'], url=r.json()[3]['store_app_url'])
        third_kb.add(button4)
    await bot.delete_message(chat_id=callback_query.from_user.id,message_id=callback_query.message.message_id)
    await bot.send_photo(callback_query.from_user.id, photo=open(str(callback_query.message.message_id) + ".jpg", "rb"),reply_markup=third_kb)
    os.remove(str(callback_query.message.message_id) + ".jpg")


if __name__ == '__main__':
    executor.start_polling(dp)