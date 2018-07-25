from flask import request
from werkzeug.security import generate_password_hash
from views import app
from model.merchant import m_register, add_commodity, de_commodity, check_commodity
from model.merchant import modify_commodity, modify_merchant


@app.route('/api/merchant/registration', methods=['post'])
def m_registration():
    data = request.get_json()
    id = data['id']
    pwd = data['password']
    password = generate_password_hash(pwd)
    email = data['email']
    addr = data['addr']
    return m_register(id, password, email, addr)


@app.route('/api/merchant/commodity', methods=['GET', 'POST'])
def commodity():
    react = request.args.get('react')
    if react == 'add':
        data = request.get_json()
        id = data['id']
        price = data['price']
        clss = data['clss']
        photo = data['photo']  # 图片传过来的是后是按照一个像素一个像素传过来的
        commodity_rest = data['commodity_rest']
        merchant_id = request.cookies.get("id")
        return add_commodity(id, price, clss, photo, commodity_rest,
                             merchant_id)
    elif react == 'delete':
        data = request.get_json()
        merchant_id = request.cookies.get("id")
        id = data['id']  # 删除商品只用传过来商品id即可
        return de_commodity(id, merchant_id)
    elif react == 'check':
        merchant_id = request.cookies.get("id")
        return check_commodity(merchant_id)


@app.route('/api/merchant/modification', methods=['post'])
def modification():
    ob = request.args.get('ob')
    data = request.get_json()
    if ob == 'merchant':
        id = data['id']
        pwd = data['password']
        password = generate_password_hash(pwd)
        email = data['email']
        addr = data['addr']
        return modify_merchant(id, password, email, addr)
    elif ob == 'commodity':
        id = data['id']
        price = data['price']
        clss = data['clss']
        photo = data['photo']  # 图片传过来的是后是按照一个像素一个像素传过来的
        commodity_rest = data['commodity_rest']
        merchant_id = request.cookies.get("id")
        return modify_commodity(id, price, clss, photo, commodity_rest,
                                merchant_id)
