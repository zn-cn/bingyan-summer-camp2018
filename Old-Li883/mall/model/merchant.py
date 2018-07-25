"""
商家操作
注册商家
注册，下架，查看物品
修改自己，物品信息
"""

from util.response import response
import pymysql

db = pymysql.connect(
    host="127.0.0.1", user="root", db="mall", passwd=',lch980929')
cursor = db.cursor()


def m_register(id, password, email, addr):
    cursor.execute("select * from merchant where id=%s", id)
    if cursor.fetchone() == (None):
        cursor.execute(
            "insert into merchant(id,password,email,identify,addr) values(%s,%s,%s,%s,%s)",
            (
                id,
                password,
                email,
                0,
                addr,
            ))
        db.commit()
        return response(200)
    else:
        return response(400, "id has been used")


def add_commodity(id, price, clss, photo, commodity_rest, merchant_id):
    cursor.execute("select * from commodity where id=%s and merchant_id=%s", (
        id,
        merchant_id,
    ))
    if cursor.fetchone() != (None):  # 判断id是否已经存在
        return response(400, "id has been used")
    f1 = open(
        '/home/oldli/weblearning/flask/project/mall/photo/' + id + '.jpg',
        'w')  # 以wb形式打开若文件不存在则会自动生成
    f1.write(photo)
    f1.close()
    photo_path = '127.0.0.1/' + id + '.jpg'  # 从前端传来文件的数据，存储在本地，前段再次调用时用nginx静态文件来获取
    cursor.execute(
        "insert into commodity(id,price,cls,photo,commodity_rest,merchant_id) values(%s,%s,%s,%s,%s,%s)",
        (
            id,
            price,
            clss,
            photo_path,
            commodity_rest,
            merchant_id,
        ))
    db.commit()
    return response(200)


def de_commodity(id, merchant_id):
    """ 
    商家下架商品
    用户收藏商品下架提醒
    """

    cursor.execute("select * from commodity where id=%s and merchant_id=%s", (
        id,
        merchant_id,
    ))
    if cursor.fetchone() != (None):
        cursor.execute("delete from commodity where id=%s and merchant_id=%s",
                       (id, merchant_id))
        db.commit()
        cursor.execute("select c_id from favourites where m_id=%s",
                       id)  # 进入该用户的消息盒子，在下一次打开该用户是发送消息给这个用户
        for i in cursor.fetchall():  # 进入每个用户的消息盒子
            cursor.execute(
                "insert into message%s(message) values('You favourite has obtained')"
                % i[0])
        return response(200)
    else:
        return response(400, "This commodity has not exist")


def check_commodity(merchant_id):
    """
    查找这个商家的所有商品
    """

    cursor.execute("select * from commodity where merchant_id=%s", merchant_id)
    commodity = []
    for t in cursor.fetchall():
        commodity.append(t)
    return response(200, commodity)


def modify_merchant(id, password, email, addr):
    """
    id作为唯一标识符，在这里不能被修改
    这里不会有用户不存在的情况，因为必须先登录才可能修改资料
    """
    cursor.execute(
        "update merchant set password=%s,email=%s,addr=%s where id=%s", (
            password,
            email,
            addr,
            id,
        ))
    db.commit()
    return response(200)


def modify_commodity(id, price, clss, photo, commodity_rest, merchant_id):
    """
    修改商品信息
    这里的商家id不变，因为商品永远属于当前商家
    """
    cursor.execute(
        "select price from commodity where id=%s and merchant_id=%s", (
            id,
            merchant_id,
        ))
    pri = cursor.fetchone()
    p = pri[0]
    if p > price:
        cursor.execute("select c_id from favourites where m_id=%s",
                       (id, ))  # 进入该用户的消息盒子，在下一次打开该用户是发送消息给这个用户
        mem = cursor.fetchall()
        for i in mem:  # 进入每个用户的消息盒子
            m = i[0]
            cursor.execute(
                "insert into message%s(message) values('You favourite thing price reduced')"
                % m)  # mysql语句中不能乱加单引号
    cursor.execute("select * from commodity where id=%s and merchant_id=%s", (
        id,
        merchant_id,
    ))
    if cursor.fetchone() == (None):  # 判断id是否存在
        return response(400, "no this modity")
    f1 = open('/home/oldli/weblearning/flask/project/mall/photo' + id + '.jpg',
              'w')  # 以wb形式打开若文件不存在则会自动生成,原来的数据会被抹去重新写
    f1.write(photo)
    cursor.execute(
        "update commodity set price=%s,cls=%s,commodity_rest=%s where id=%s", (
            price,
            clss,
            commodity_rest,
            id,
        ))
    db.commit()
    return response(200)
