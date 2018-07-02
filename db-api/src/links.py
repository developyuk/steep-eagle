def links_get_id(data):
  if type(data) is dict:
    return data['id']
  else:
    return data


def links_users(data):
  if 'channels' in data:
    data['_links']['channels'] = list(
      map(lambda v: {'title': 'channels', 'href': 'channels/%s' % links_get_id(v)}, data['channels']))
  return data


def links_channels(data):
  if 'owner' in data:
    data['_links']['owner'] = {'title': 'owner', 'href': 'users/%s' % links_get_id(data['owner'])}
  if 'courses' in data:
    data['_links']['courses'] = list(
      map(lambda v: {'title': 'courses', 'href': 'courses/%s' % links_get_id(v)}, data['courses']))
  if 'events' in data:
    data['_links']['events'] = list(
      map(lambda v: {'title': 'events', 'href': 'events/%s' % links_get_id(v)}, data['events']))

  return data


def links_courses(data):
  if 'channel' in data:
    data['_links']['channel'] = {'title': 'channel', 'href': 'channels/%s' % links_get_id(data['channel'])}
  if 'lessons' in data:
    data['_links']['lessons'] = list(
      map(lambda v: {'title': 'lessons', 'href': 'lessons/%s' % links_get_id(v)}, data['lessons']))
  return data


def links_events(data):
  data = links_courses(data)
  if 'location' in data:
    data['_links']['location'] = {'title': 'location', 'href': 'locations/%s' % links_get_id(data['location'])}
  if 'sessions' in data:
    data['_links']['sessions'] = list(
      map(lambda v: {'title': 'sessions', 'href': 'sessions/%s' % links_get_id(v)}, data['sessions']))
  return data


def links_lessons(data):
  if 'course' in data:
    data['_links']['course'] = {'title': 'course', 'href': 'course/%s' % links_get_id(data['course'])}
  return data


def links_locations(data):
  if 'event' in data:
    data['_links']['event'] = {'title': 'event', 'href': 'events/%s' % links_get_id(data['event'])}
  return data


def links_sessions(data):
  if 'event' in data:
    data['_links']['event'] = {'title': 'event', 'href': 'events/%s' % links_get_id(data['event'])}
  return data


def before_returning_items(resource_name, response):
  print('About to return items from "%s" ' % resource_name)

  if resource_name == 'users':
    response['_items'] = list(map(links_users, response['_items']))

  if resource_name == 'channels':
    response['_items'] = list(map(links_channels, response['_items']))

  if resource_name == 'courses':
    response['_items'] = list(map(links_courses, response['_items']))

  if resource_name == 'lessons':
    response['_items'] = list(map(links_lessons, response['_items']))

  if resource_name == 'events':
    response['_items'] = list(map(links_events, response['_items']))

  if resource_name == 'locations':
    response['_items'] = list(map(links_locations, response['_items']))

  if resource_name == 'sessions':
    response['_items'] = list(map(links_sessions, response['_items']))


def before_returning_item(resource_name, response):
  print('About to return an item from "%s" ' % resource_name)

  if resource_name == 'users':
    response = links_users(response)

  if resource_name == 'channels':
    response = links_channels(response)

  if resource_name == 'courses':
    response = links_courses(response)

  if resource_name == 'lessons':
    response = links_lessons(response)

  if resource_name == 'events':
    response = links_events(response)

  if resource_name == 'locations':
    response = links_locations(response)

  if resource_name == 'sessions':
    response = links_sessions(response)
