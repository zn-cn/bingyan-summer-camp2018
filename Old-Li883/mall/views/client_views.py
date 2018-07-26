from flask import request
from werkzeug.security import generate_password_hash
from views import app
from model.client import register, find_commodity_addr, find_commodity_cls
from model.client import find_info, modify_info
from model.client import in_shopping_cart, out_shopping_cart, check_shopping_cart
from model.client import in_favourites, out_favourites, check_favourites, check_commodity
from model.client import buy_thing
from model.client import add_addr, modify_addr, check_addr
from model.client import find_message


@app.route('/api/client/registration', methods=['post'])
def registration():
    """
    注册时除基本信息外还要写清楚收货地址
    """

    data = request.get_json()
    id = data['id']
    pwd = data['password']
    password = generate_password_hash(pwd)
    email = data['email']
    name = data['name']
    addr = data['addr']
    phone = data['phone']
    return register(id, password, email, name, addr, phone)


@app.route('/api/client/commodity/location', methods=['post'])
def commodity_location():
    """
    按分类查询商品，用query表名查询方式及类型
    """

    way = request.args.get('way')
    data = request.get_json()
    name = data['name']
    if way == 'cls':
        return find_commodity_cls(name)
    elif way == 'addr':
        return find_commodity_addr(name)


@app.route('/api/client/info', methods=['post', 'get'])
def info():
    operate = request.args.get('operate')
    if operate == 'find':
        id = request.cookies.get('id')
        data = request.get_json()
        who = data['who']  # 查自己的还是商家的信息
        return find_info(who, id)
    elif operate == 'modify':
        data = request.get_json()
        id = request.cookies.get("id")
        pwd = data['password']
        password = generate_password_hash(pwd)
        email = data['email']
        return modify_info(id, password, email)


@app.route('/api/client/cart', methods=['post', 'get'])
def cart():
    operate = request.args.get('operate')
    c_id = request.cookies.get("id")
    if operate == 'in':
        data = request.get_json()
        id = data['id']  # 这里是物品id
        return in_shopping_cart(c_id, id)
    elif operate == 'out':
        data = request.get_json()
        id = data['id']  # 这里是物品id
        return out_shopping_cart(c_id, id)
    elif operate == 'check':
        return check_shopping_cart(c_id)


@app.route('/api/client/favourites', methods=['post', 'get'])
def favourites():
    """
    所有用户的收藏夹用一个数据库
    方便后面的查找
    """

    operate = request.args.get('operate')
    c_id = request.cookies.get("id")
    if operate == 'in':
        data = request.get_json()
        m_id = data['id']  # 这里是物品id
        return in_favourites(c_id, m_id)
    elif operate == 'out':
        data = request.get_json()
        m_id = data['id']  # 这里是物品id
        return out_favourites(c_id, m_id)
    elif operate == 'check':
        return check_favourites(c_id)


@app.route('/api/client/commodity')
def check_commodity_info():
    c_id = request.cookies.get("id")
    id = request.args.get("id")
    return check_commodity(c_id, id)


@app.route('/api/client/buy', methods=['post'])
def buy():
    c_id = request.cookies.get("id")
    data = request.get_json()
    id = data['id']
    return buy_thing(c_id, id)


@app.route('/api/client/addr', methods=['post', 'get'])
def addr():
    c_id = request.cookies.get("id")
    operate = request.args.get('operate')
    if operate == 'add':
        data = request.get_json()
        name = data['name']
        ad = data['ad']
        phone = data['phone']
        return add_addr(c_id, name, ad, phone)
    elif operate == 'modify':
        data = request.get_json()
        name = data['name']
        ad = data['ad']
        phone = data['phone']
        return modify_addr(c_id, name, ad, phone)
    elif operate == 'check':
        check_addr(c_id)


@app.route('/api/client/message')
def message():
    id = request.cookies.get("id")
    return find_message(id)
