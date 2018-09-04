from copy import deepcopy

from flask import current_app as app
from werkzeug import secure_filename
from tables import ClassStudents


def fclass_students(id, students_):
    session = app.data.driver.session

    session.query(ClassStudents).filter_by(class_id=id).delete()
    session.flush()

    class_students = map(lambda v: ClassStudents(
        class_id=id, student_id=v), students_)
    session.add_all(class_students)
    session.commit()


def on_update(updates, original):
    students_ = updates.get('students_')

    fclass_students(original.get('id'), students_)


post_items = []


def on_insert(items):
    global post_items
    post_items = deepcopy(items)
    for i, item in enumerate(items):
        del item['students_']


def on_inserted(items):
    global post_items
    for i, item in enumerate(items):
        students_ = post_items[i].get('students_')
        if students_:
            fclass_students(item.get('id'), students_)
    post_items = []
