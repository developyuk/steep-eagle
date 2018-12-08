user = {
    'username': {
        'type': 'string',
        'unique': True,
    },
    'password': {
        'type': 'string',
    },
    'role': {
        'type': 'string',
        'required': True,
        'allowed': ['tutor', 'operation', 'student', 'guardian']
    },
    'email': {
        'type': 'string',
        'unique': True,
    },
    'name': {
        'type': 'string',
    },
    'dob': {
        'type': 'datetime',
    },
    'photo': {
        'type': 'media',
    },
    'address': {
        'type': 'string',
    },
    'school': {
        'type': 'string',
    },
    'grade': {
        'type': 'string',
    },
    'contact': {
        'type': 'string',
    },
    'student': {
        'type': 'integer',
        'data_relation': {
            'resource': 'users',
            'field': '_id',
            'embeddable': True
        },
    },
}

branch = {
    'name': {
        'type': 'string',
        'required': True
    },
    'address': {
        'type': 'string',
    },
}

module = {
    'name': {
        'type': 'string',
        'required': True
    },
    'image': {
        'type': 'media',
    },
}

class_ = {
    'start': {
        'type': 'datetime'
    },
    'finish': {
        'type': 'datetime'
    },
    'schedule': {
        'type': 'dict',
        'schema': {
            'recurrence': {
                'type': 'dict',
                'schema': {
                    'interval': {
                        'type': 'integer'
                    },
                    'freq': {
                        'type': 'string',
                        'allowed': ['daily', 'weekly', 'monthly', 'yearly']
                    },
                    'byday': {
                        'type': 'list',
                        'allowed': ['sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday']
                    },
                    'until': {
                        'type': 'datetime'
                    },
                    'count': {
                        'type': 'integer'
                    }
                }
            },
            'exclude': {
                'type': 'list',
            },
            'include': {
                'type': 'list',
            },
        }
    },
    'module': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'modules',
            'field': '_id',
            'embeddable': True
        },
    },
    'branch': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'branches',
            'field': '_id',
            'embeddable': True
        },
    },
    'tutor': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'users',
            'field': '_id',
            'embeddable': True
        },
    },
    'studentsTotal': {
        'type': 'integer',
        'default': 0
    },
}

schedules = {
    'class': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'classes',
            'field': '_id',
            'embeddable': True
        },
    },
    'start': {
        'type': 'datetime',
        'required': True
    },
    'finish': {
        'type': 'datetime',
        'required': True
    },
}

class_student = {
    'class': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'classes',
            'field': '_id',
            'embeddable': True
        },
    },
    'student': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'users',
            'field': '_id',
            'embeddable': True
        },
    },
}

attendance = {
    'class': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'classes',
            'field': '_id',
            'embeddable': True
        },
    },
    'module': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'modules',
            'field': '_id',
            'embeddable': True
        },
    },
}

attendance_tutor = {
    'attendance': {
        'type': 'integer',
        'data_relation': {
            'resource': 'attendances',
            'field': '_id',
            'embeddable': True
        },
    },
    'tutor': {
        'type': 'integer',
        'data_relation': {
            'resource': 'users',
            'field': '_id',
            'embeddable': True
        },
    },
}
attendance_student = {
    'attendance': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'attendances',
            'field': '_id',
            'embeddable': True
        },
    },
    'tutor': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'users',
            'field': '_id',
            'embeddable': True
        },
    },
    'student': {
        'type': 'integer',
        'required': True,
        'data_relation': {
            'resource': 'users',
            'field': '_id',
            'embeddable': True
        },
    },
    'isPresence': {
        'type': 'boolean',
        'default': False
    },
    'feedback': {
        'type': 'string',
    },
    'rating': {
        'type': 'dict',
        'schema': {
            'interaction': {'type': 'integer'},
            'cognition': {'type': 'integer'},
            'creativity': {'type': 'integer'},
        },
    },
}
cache = {
    'key': {
        'type': 'string',
    },
    'value': {
        'type': 'dict',
    },
}
