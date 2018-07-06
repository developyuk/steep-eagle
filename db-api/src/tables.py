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
  pass_ = Column('pass', String)
  profile = relationship("UsersProfile", uselist=False, back_populates="user")

  @staticmethod
  def auth(token):
    try:
      data = jwt.decode(token, app.config['JWT_SECRET'])
    except Exception as e:
      return jsonify({'message':str(e)}), 400

    return data

  def sign(self):
    data = {
      'id': self.id,
    }
    return jwt.encode(data, app.config['JWT_SECRET'])


class UsersProfile(Base):
  __tablename__ = 'users_profile'
  name = Column(String)
  dob = Column(String)
  photo = Column(String)
  user_id = Column(Integer, ForeignKey('users.id'), primary_key=True)
  user = relationship("Users", back_populates="profile")


class Branches(CommonColumns):
  __tablename__ = 'branches'
  id = Column(Integer, primary_key=True, autoincrement=True)
  name = Column(String)
  address = Column(String)


class ProgramTypes(CommonColumns):
  __tablename__ = 'program_types'
  id = Column(Integer, primary_key=True, autoincrement=True)
  name = Column(String)
  programs = relationship("Programs", back_populates="type")


class ProgramsModules(Base):
  __tablename__ = 'programs_modules'
  id = Column(Integer, primary_key=True, autoincrement=True)
  program_id = Column(Integer, ForeignKey('programs.id'))
  program = relationship("Programs", back_populates="modules")
  module_id = Column(Integer, ForeignKey('modules.id'))
  module = relationship("Modules", back_populates="programs")
  # class_ = relationship("Classes", back_populates="program_module")


class Programs(CommonColumns):
  __tablename__ = 'programs'
  id = Column(Integer, primary_key=True, autoincrement=True)
  name = Column(String)
  type_id = Column(Integer, ForeignKey('program_types.id'))
  type = relationship("ProgramTypes", back_populates="programs")
  modules = relationship("ProgramsModules", back_populates="program")


class Modules(CommonColumns):
  __tablename__ = 'modules'
  id = Column(Integer, primary_key=True, autoincrement=True)
  name = Column(String)
  image = Column(String)
  programs = relationship("ProgramsModules", back_populates="module")


class ClassStudents(Base):
  __tablename__ = 'class_students'
  id = Column(Integer, primary_key=True, autoincrement=True)
  class_id = Column(Integer, ForeignKey('classes.id'))
  class_ = relationship("Classes")
  student_id = Column(Integer, ForeignKey('users.id'))
  student = relationship("Users")


class Classes(CommonColumns):
  __tablename__ = 'classes'
  id = Column(Integer, primary_key=True, autoincrement=True)
  day = Column(String)
  start_at = Column(String)
  finish_at = Column(String)
  program_module_id = Column(Integer, ForeignKey('programs_modules.id'))
  program_module = relationship("ProgramsModules")
  branch_id = Column(Integer, ForeignKey('branches.id'))
  branch = relationship("Branches")
  tutor_id = Column(Integer, ForeignKey('users.id'))
  tutor = relationship("Users")
  students = relationship("ClassStudents")
  # sessions = relationship("Sessions", back_populates="class_")


class Sessions(CommonColumns):
  __tablename__ = 'sessions'
  id = Column(Integer, primary_key=True, autoincrement=True)
  class_id = Column(Integer, ForeignKey('classes.id'))
  class_ = relationship("Classes")
  session_tutors = relationship("SessionsTutors", back_populates="session")


class SessionsTutors(CommonColumns):
  __tablename__ = 'sessions_tutors'
  id = Column(Integer, primary_key=True, autoincrement=True)
  session_id = Column(Integer, ForeignKey('sessions.id'))
  session = relationship("Sessions", back_populates="session_tutors")
  tutor_id = Column(Integer, ForeignKey('users.id'))
  tutor = relationship("Users")
  session_students = relationship("SessionsTutorsStudents", back_populates="session_tutor")


class SessionsTutorsStudents(CommonColumns):
  __tablename__ = 'sessions_tutors_students'
  id = Column(Integer, primary_key=True, autoincrement=True)
  status = Column(Boolean)
  feedback = Column(String)
  rating_interaction = Column(Integer)
  rating_cognition = Column(Integer)
  rating_creativity = Column(Integer)
  session_tutor_id = Column(Integer, ForeignKey('sessions_tutors.id'))
  session_tutor = relationship("SessionsTutors", back_populates="session_students")
  student_id = Column(Integer, ForeignKey('users.id'))
  student = relationship("Users")

dow = dict(zip('monday tuesday wednesday thursday friday saturday sunday'.split(), range(7)))

class ClassesTs(Classes):
  def onDay(self):
    now = datetime.now()
    return now + timedelta(days=(dow[self.day] - now.weekday() + 7) % 7)

  @hybrid_property
  def start_at_ts(self):
    dt = datetime.strptime('%sT%s' % (self.onDay().date(), self.start_at), '%Y-%m-%dT%H:%M')
    return timezone('Asia/Jakarta').localize(dt).astimezone(timezone('UTC'))

  @start_at_ts.expression
  def start_at_ts(cls):
    return cls

  @hybrid_property
  def finish_at_ts(self):
    dt = datetime.strptime('%sT%s' % (self.onDay().date(), self.finish_at), '%Y-%m-%dT%H:%M')
    return timezone('Asia/Jakarta').localize(dt).astimezone(timezone('UTC'))

  @finish_at_ts.expression
  def finish_at_ts(cls):
    return cls

  @hybrid_property
  def q(self):
    return ', '.join((self.program_module.module.name, self.branch.name, self.day, self.start_at, self.finish_at,
                      self.tutor.username, self.tutor.profile.name))

  @q.expression
  def q(cls):
    return cls


class Exports(ClassStudents):
  @hybrid_property
  def program_name(self):
    return self.class_.program_module.program.type.name

  @program_name.expression
  def program_name(cls):
    return Classes.program_module

  @hybrid_property
  def program(self):
    return self.class_.program_module.program.name

  @program.expression
  def program(cls):
    return Classes.program_module

  @hybrid_property
  def module(self):
    return self.class_.program_module.module.name

  @module.expression
  def module(cls):
    return Classes.program_module

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
    return self.student.profile.name

  @student_name.expression
  def student_name(cls):
    return Users.profile

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
    return self.class_.tutor.profile.name

  @tutor_name.expression
  def tutor_name(cls):
    return Classes.tutor
