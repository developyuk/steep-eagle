import os
from pprint import pprint

from flask import current_app as app, jsonify
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
    class_students = map(lambda v: v.id, class_students)

    return class_students


def before_patch_item(updates, original):
    students_ = updates.get('students_')
    if students_:
        fclass_students(original.get('id'), students_)


def after_post_item(items):
    for i, item in enumerate(items):
        students_ = item.get('students_')
        if students_:
            fclass_students(item.get('id'), students_)
        #     items[i].update({'students': students})
