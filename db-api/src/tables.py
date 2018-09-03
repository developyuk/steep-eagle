
from datetime import datetime, timedelta

# from flask import current_app as app, jsonify
from pytz import timezone

from sqlalchemy import Column, DateTime, ForeignKey, Integer, String, Boolean, Text, func, select
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.ext.hybrid import hybrid_property
from sqlalchemy.orm import column_property, relationship
from blueprints import dow

Base = declarative_base()


class CommonColumns(Base):
    __abstract__ = True
    _created = Column(DateTime, default=func.now())
    _updated = Column(DateTime, default=func.now(), onupdate=func.now())
    _etag = Column(String(40))


class Users(CommonColumns):
    __tablename__ = 'users'
    id = Column(Integer, primary_key=True, autoincrement=True)
    username = Column(String)
    email = Column(String)
    role = Column(String)
    name = Column(String)
    dob = Column(String)
    photo = Column(String)
    pass_ = Column('pass', String)

    _deleted = Column(Boolean, default=False)


class Branches(CommonColumns):
    __tablename__ = 'branches'
    id = Column(Integer, primary_key=True, autoincrement=True)
    name = Column(String)
    address = Column(String)


class Modules(CommonColumns):
    __tablename__ = 'modules'
    id = Column(Integer, primary_key=True, autoincrement=True)
    name = Column(String)
    image = Column(String)


class ClassStudents(CommonColumns):
    __tablename__ = 'class_students'
    id = Column(Integer, primary_key=True, autoincrement=True)
    class_id = Column(Integer, ForeignKey('classes.id'))
    student_id = Column(Integer, ForeignKey('users.id'))

    class_ = relationship("ClassesTs")
    student = relationship("Users")


class Classes(CommonColumns):
    __tablename__ = 'classes'
    id = Column(Integer, primary_key=True, autoincrement=True)
    day = Column(String)
    start_at = Column(String)
    finish_at = Column(String)
    module_id = Column(Integer, ForeignKey('modules.id'))
    branch_id = Column(Integer, ForeignKey('branches.id'))
    tutor_id = Column(Integer, ForeignKey('users.id'))

    module = relationship("Modules")
    branch = relationship("Branches")
    tutor = relationship("Users")
    students = relationship("ClassStudents")
    attendances = relationship("Attendances", back_populates="class_")

    q = column_property(
        select([Modules.name+" " + Branches.name+" "+day+" "+start_at+" "+finish_at+" "+Users.username+" "+Users.name]).
        where(Branches.id == branch_id).
        where(Modules.id == module_id).
        where(Users.id == tutor_id)
    )


class Attendances(CommonColumns):
    __tablename__ = 'attendances'
    id = Column(Integer, primary_key=True, autoincrement=True)
    class_id = Column(Integer, ForeignKey('classes.id'))
    module_id = Column(Integer, ForeignKey('modules.id'))

    class_ = relationship("ClassesTs", back_populates="attendances")
    module = relationship("Modules")
    attendance_tutors = relationship(
        "AttendancesTutors", back_populates="attendance")
    attendance_students = relationship(
        "AttendancesStudents", back_populates="attendance")


class AttendancesTutors(CommonColumns):
    __tablename__ = 'attendances_tutors'
    id = Column(Integer, primary_key=True, autoincrement=True)
    attendance_id = Column(Integer, ForeignKey('attendances.id'))
    tutor_id = Column(Integer, ForeignKey('users.id'))
    tutor = relationship("Users")

    attendance = relationship(
        "Attendances", back_populates="attendance_tutors")
    # attendance_students = relationship("AttendancesTutorsStudents", back_populates="attendance_tutor")


class AttendancesStudents(CommonColumns):
    __tablename__ = 'attendances_students'
    id = Column(Integer, primary_key=True, autoincrement=True)
    attendance_id = Column(Integer, ForeignKey('attendances.id'))
    tutor_id = Column(Integer, ForeignKey('users.id'))
    student_id = Column(Integer, ForeignKey('users.id'))
    is_presence = Column(Boolean)
    feedback = Column(String)
    rating_interaction = Column(Integer)
    rating_cognition = Column(Integer)
    rating_creativity = Column(Integer)
    # attendance_tutor_id = Column(Integer, ForeignKey('attendances_tutors.id'))
    # attendance_tutor = relationship( "AttendancesTutors", back_populates="attendance_students")

    attendance = relationship(
        "Attendances", back_populates="attendance_students")
    student = relationship("Users", foreign_keys=[student_id])
    tutor = relationship("Users", foreign_keys=[tutor_id])


def onDay(day):
    now = datetime.now()
    return now + timedelta(days=(dow[day] - now.weekday() + 7) % 7)


class ClassesTs(Classes):

    @hybrid_property
    def start_at_ts(self):
        dt = datetime.strptime('%sT%s' % (
            onDay(self.day).date(), self.start_at), '%Y-%m-%dT%H:%M')
        return timezone('Asia/Jakarta').localize(dt)

    @start_at_ts.expression
    def start_at_ts(cls):
        return cls

    @hybrid_property
    def finish_at_ts(self):
        dt = datetime.strptime('%sT%s' % (
            onDay(self.day).date(), self.finish_at), '%Y-%m-%dT%H:%M')
        return timezone('Asia/Jakarta').localize(dt)

    @finish_at_ts.expression
    def finish_at_ts(cls):
        return cls


class Notifications(CommonColumns):
    __tablename__ = 'notifications'
    id = Column(Integer, primary_key=True, autoincrement=True)
    message = Column(String)
    is_read = Column(Boolean)
    data = Column(String)
