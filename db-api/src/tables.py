from sqlalchemy import Table, Column, DateTime, ForeignKey, Integer, String, Boolean, func
# from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.ext.hybrid import hybrid_property, hybrid_method
from sqlalchemy.orm import column_property, relationship
from sqlservice import declarative_base
import jwt
from pprint import pprint
from datetime import datetime, timedelta
from flask import current_app as app, jsonify
from pytz import timezone
# from eve.methods.common import embedded_document

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

    @staticmethod
    def auth(token):
        try:
            data = jwt.decode(token, app.config['JWT_SECRET'])
        except Exception as e:
            return jsonify({'message': str(e)}), 400

        return data

    def sign(self):
        data = {
            'id': self.id,
        }
        return jwt.encode(data, app.config['JWT_SECRET'])


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


class ClassStudents(Base):
    __tablename__ = 'class_students'
    id = Column(Integer, primary_key=True, autoincrement=True)
    class_id = Column(Integer, ForeignKey('classes.id'))
    class_ = relationship("ClassesTs")
    student_id = Column(Integer, ForeignKey('users.id'))
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
    # sessions = relationship("Sessions", back_populates="class_")


class Sessions(CommonColumns):
    __tablename__ = 'sessions'
    id = Column(Integer, primary_key=True, autoincrement=True)
    class_id = Column(Integer, ForeignKey('classes.id'))
    class_ = relationship("ClassesTs")
    session_tutors = relationship("SessionsTutors", back_populates="session")


class SessionsTutors(CommonColumns):
    __tablename__ = 'sessions_tutors'
    id = Column(Integer, primary_key=True, autoincrement=True)
    session_id = Column(Integer, ForeignKey('sessions.id'))
    session = relationship("Sessions", back_populates="session_tutors")
    tutor_id = Column(Integer, ForeignKey('users.id'))
    tutor = relationship("Users")
    session_students = relationship(
        "SessionsTutorsStudents", back_populates="session_tutor")


class SessionsTutorsStudents(CommonColumns):
    __tablename__ = 'sessions_tutors_students'
    id = Column(Integer, primary_key=True, autoincrement=True)
    status = Column(Boolean)
    feedback = Column(String)
    rating_interaction = Column(Integer)
    rating_cognition = Column(Integer)
    rating_creativity = Column(Integer)
    session_tutor_id = Column(Integer, ForeignKey('sessions_tutors.id'))
    session_tutor = relationship(
        "SessionsTutors", back_populates="session_students")
    student_id = Column(Integer, ForeignKey('users.id'))
    student = relationship("Users")


dow = dict(
    zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))


def onDay(day):
    now = datetime.now()
    return now + timedelta(days=(dow[day] - now.weekday() + 7) % 7)


class ClassesTs(Classes):
    @hybrid_property
    def start_at_ts(self):
        dt = datetime.strptime('%sT%s' % (
            onDay(self.day).date(), self.start_at), '%Y-%m-%dT%H:%M')
        return timezone('Asia/Jakarta').localize(dt).astimezone(timezone('UTC'))

    @start_at_ts.expression
    def start_at_ts(cls):
        return cls

    @hybrid_property
    def finish_at_ts(self):
        dt = datetime.strptime('%sT%s' % (
            onDay(self.day).date(), self.finish_at), '%Y-%m-%dT%H:%M')
        return timezone('Asia/Jakarta').localize(dt).astimezone(timezone('UTC'))

    @finish_at_ts.expression
    def finish_at_ts(cls):
        return cls

    @hybrid_property
    def q(self):
        return ', '.join((self.module.name, self.branch.name, self.day, self.start_at, self.finish_at,
                          self.tutor.username, self.tutor.name))

    @q.expression
    def q(cls):
        return cls


class Exports(ClassStudents):
    @hybrid_property
    def branch(self):
        return self.class_.branch.name

    @branch.expression
    def branch(cls):
        return Classes.branch

    @hybrid_property
    def day(self):
        return self.class_.day

    @day.expression
    def day(cls):
        return Classes.day

    @hybrid_property
    def start_at(self):
        return self.class_.start_at

    @start_at.expression
    def start_at(cls):
        return Classes.start_at

    @hybrid_property
    def finish_at(self):
        return self.class_.finish_at

    @finish_at.expression
    def finish_at(cls):
        return Classes.finish_at

    @hybrid_property
    def student_name(self):
        return self.student.name

    @student_name.expression
    def student_name(cls):
        return Users

    @hybrid_property
    def tutor_email(self):
        return self.class_.tutor.email

    @tutor_email.expression
    def tutor_email(cls):
        return Classes.tutor

    @hybrid_property
    def tutor_username(self):
        return self.class_.tutor.username

    @tutor_username.expression
    def tutor_username(cls):
        return Classes.tutor

    @hybrid_property
    def tutor_name(self):
        return self.class_.tutor.name

    @tutor_name.expression
    def tutor_name(cls):
        return Classes.tutor
