from sqlalchemy import Table, Column, DateTime, ForeignKey, Integer, String, Boolean, func
from sqlalchemy.ext.declarative import declarative_base

from sqlalchemy.orm import column_property, relationship
import jwt
from flask import current_app as app

Base = declarative_base()


class CommonColumns(Base):
  __abstract__ = True
  _created = Column(DateTime, default=func.now())
  _updated = Column(DateTime, default=func.now(), onupdate=func.now())
  # _etag = Column(String(40))


class Users(CommonColumns):
  __tablename__ = 'users'
  id = Column(Integer, primary_key=True, autoincrement=True)
  username = Column(String)
  email = Column(String)
  role = Column(String)
  pass_ = Column('pass', String)
  # classes = relationship("Classes", back_populates="tutor")
  profile = relationship("UsersProfile", uselist=False, back_populates="user")

  @staticmethod
  def auth(token):
    try:
      data = jwt.decode(token, app.config['JWT_SECRET'])
    except jwt.ExpiredSignatureError:
      # Signature has expired
      return None
    except jwt.DecodeError:
      # Signature has expired
      return None

    return data

  def sign(self):
    return jwt.encode({'username': self.username, 'id': self.id}, app.config['JWT_SECRET'])


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
