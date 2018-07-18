from flask import make_response
import json
import pymysql

db = pymysql.connect(
    host="127.0.0.1", user="root", db="member", passwd=',lch980929')
cursor = db.cursor()


def mem_login(id, password):
    cursor.execute("select password from mem_info where id=%s", id)
    accurate_pwd = cursor.fetchone()
    if accurate_pwd:
        if accurate_pwd[0] == password:
            rsp = make_response('hello')
            rsp.set_cookie('id', id)  # 返回登录者id
            cursor.execute("select status from mem_info where id=%s", id)
            now_status = cursor.fetchone()
            if now_status[0] == 1 and now_status[0] != 0:  # 普通用户登录
                cursor.execute("update mem_info set status=3 where id=%s", id)
            elif now_status[0] == 2 and now_status[0] != 0:  # 管理员登录
                cursor.execute("update mem_info set status=4 where id=%s", id)
            elif now_status[0] > 2:  # 判断是否已经登录
                return "you hava logined"
            elif now_status[0] == 0:  # 还未被认证的用户登录
                return "you have not been identified"
            db.commit()
            return rsp
    return 'id or password wrong'


def mem_register(id, password, name, email, groups):
    cursor.execute("select * from mem_info where id=%s||name=%s", (
        id,
        name,
    ))
    if cursor.fetchone() == (None):
        cursor.execute(
            "insert into mem_info(id,name,password,email,groups,status) values(%s,%s,%s,%s,%s,0)",
            (
                id,
                name,
                password,
                email,
                groups,
            ))
        db.commit()
        return "you success register,wait for identify"
    else:
        return "the id or name have used"


def find_no_register():
    cursor.execute("select * from mem_info where status=0")
    no_register_mem = []
    find_mem = cursor.fetchall()
    for mem in find_mem:
        no_register_mem.append(mem)
    return json.dumps(no_register_mem)


def add_new_member(id, password, name, email, groups, status):
    cursor.execute("select * from mem_info where id=%s||name=%s", (
        id,
        name,
    ))
    if cursor.fetchone() == (None):
        if status == 1:
            cursor.execute(
                "insert into mem_info(id,name,password,email,groups,status) values(%s,%s,%s,%s,%s,1)",
                (
                    id,
                    name,
                    password,
                    email,
                    groups,
                ))
        elif status == 2:
            cursor.execute(
                "insert into mem_info(id,name,password,email,groups,status) values(%s,%s,%s,%s,%s,2)",
                (
                    id,
                    name,
                    password,
                    email,
                    groups,
                ))
        db.commit()
        return "you success add"
    else:
        return "the id or name have used"


def delete_member(id):
    cursor.execute("select * from mem_info where id=%s", id)
    if cursor.fetchone() == (None):
        return "no this member"
    cursor.execute("delete from mem_info where id=%s", id)
    db.commit()
    return "success delete this member"


def change_member(id, password, name, email, groups):
    if password:
        cursor.execute("update mem_info set password=%s where id=%s", (
            password,
            id,
        ))
    if name:
        cursor.execute("update mem_info set name=%s where id=%s", (
            name,
            id,
        ))
    if email:
        cursor.execute("update mem_info set email=%s where id=%s", (
            email,
            id,
        ))
    if groups:
        cursor.execute("update mem_info set groups=%s where id=%s", (
            groups,
            id,
        ))
    db.commit()
    return "you have successfully change this member's information"


def find_all():
    cursor.execute("select * from mem_info")
    members = []
    find_mem = cursor.fetchall()
    for mem in find_mem:
        members.append(mem)
    return json.dumps(members)


def group_member(groups):
    cursor.execute("select * from mem_info where groups=%s", groups)
    members = []
    find_mem = cursor.fetchall()
    i = 0
    for mem in find_mem:
        if find_mem[i][5] != 0:  # 未认证的不算组员
            members.append(mem)
        i += 1
    return json.dumps(members)


def mem_logout(id):
    cursor.execute("select status from mem_info where id=%s", id)
    now_status = cursor.fetchone()
    if now_status[0] == 3:  # 普通用户登出
        cursor.execute("update mem_info set status=1 where id=%s", id)
    elif now_status[0] == 4:  # 管理员登出
        cursor.execute("update mem_info set status=2 where id=%s", id)
    resp = make_response("you have successfully logout")
    resp.delete_cookie('id')
    db.commit()
    return resp
