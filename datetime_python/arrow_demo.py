import arrow
from dateutil import tz
from datetime import datetime

def construct_demo():
    """
    Construct arrow object
    """
    # get now of local time, timezone from system, AEST in my case
    dt = arrow.now() 
    print(dt.isoformat())       # 2021-06-18T07:35:21.865380+10:00
    # get now of UTC time
    dt = arrow.utcnow()
    print(dt.isoformat()) 
    # get now of another timezone
    dt = arrow.now('US/Pacific')
    # get from ISO 8601 string
    dt = arrow.get('2013-05-11T21:23:58.970460+07:00')
    # get as datetime, UTC
    dt = arrow.get(2013, 5, 5)  # 2013-05-05T00:00:00+00:00
    # get from UNIX time, to UTC
    dt = arrow.get(1623966051)
    print(dt.isoformat())       # 2021-06-17T21:40:51+00:00, 默认为Timezone为UTC
    # get from UNIX time, to local(AEST in my case)
    dt = arrow.get(1623966051, tzinfo='local')
    # get from UNIX time, to another timezone, support IANA and abbreviation
    dt = arrow.get(1623966051, tzinfo='AEST')
    print(dt.isoformat())       # 2021-06-18T07:40:51+10:00
    dt = arrow.get(1623966051, tzinfo='Australia/Sydney')
    print(dt.isoformat())       # 2021-06-18T07:40:51+10:00

def timezone_convert_demo():
    dt = arrow.now() 
    print(dt.to('UTC').isoformat())                 # 2021-06-17T21:42:46.497456+00:00
    print(dt.to('Australia/Sydney').isoformat())    # 2021-06-18T07:42:46.497456+10:00
    dt.to(tz.gettz('US/Pacific')) # support Python tz package

def parse_format_demo():
    # parse from own format, default UTC
    dt = arrow.get('2013-05-05 12:30:45', 'YYYY-MM-DD HH:mm:ss')
    print(dt.isoformat())                           # 2013-05-05T12:30:45+00:00
    print(dt.format('YYYY-MM-DD HH:mm:ss ZZ') )     # 2013-05-07 05:23:16 -00:00, ZZ代表offset

def properties_demo():
    dt = arrow.utcnow()
    # get Python's datetime.datetime
    dt.datetime     # datetime.datetime(2013, 5, 7, 4, 38, 15, 447644, tzinfo=tzutc())
    dt.tzinfo       # tzutc()
    dt.year
    dt.date()       # datetime.date(2013, 5, 7)
    dt.time()       # datetime.time(4, 38, 15, 447644)
    dt.timetz()     # datetime.time(22, 4, 50, 202752, tzinfo=tzutc())

def replace_shift_demo():
    # replace some fields, not change other attribute
    arw = arrow.utcnow()                    # 2021-06-17T08:12:37.986011+00:00
    arw = arw.replace(hour=4, minute=40)    # 2021-06-17T04:40:37.986011+00:00
    arw = arw.shift(weeks=+3)               # 2021-07-08T04:40:37.986011+00:00
    # only change timezone, not update other fields, different with convert
    arw = arw.replace(tzinfo='US/Pacific')  # 2021-07-08T04:40:37.986011-07:0

def range_span_demo():
    # 当前time的hour开始及结束, 当前时间: 2021-06-17T22:15:38.970255+00:00
    arrow.utcnow().span('hour') # (<Arrow [2021-06-17T22:00:00+00:00]>, <Arrow [2021-06-17T22:59:59.999999+00:00]>)
    # 另一种方法是floor和ceiling, 得出一样的span start和end
    hour_start = arrow.utcnow().floor('hour')
    hour_end = arrow.utcnow().ceil('hour')
    # get of range of time spans
    start = datetime(2013, 5, 5, 12, 30)
    end = datetime(2013, 5, 5, 17, 15)
    for r in arrow.Arrow.span_range('hour', start, end):
        print(r)
    """
    (<Arrow [2013-05-05T12:00:00+00:00]>, <Arrow [2013-05-05T12:59:59.999999+00:00]>)
    (<Arrow [2013-05-05T13:00:00+00:00]>, <Arrow [2013-05-05T13:59:59.999999+00:00]>)
    (<Arrow [2013-05-05T14:00:00+00:00]>, <Arrow [2013-05-05T14:59:59.999999+00:00]>)
    (<Arrow [2013-05-05T15:00:00+00:00]>, <Arrow [2013-05-05T15:59:59.999999+00:00]>)
    (<Arrow [2013-05-05T16:00:00+00:00]>, <Arrow [2013-05-05T16:59:59.999999+00:00]>)
    """
    # iterate over range of times
    for r in arrow.Arrow.range('hour', start, end):
        print(repr(r))
    """
    <Arrow [2013-05-05T12:30:00+00:00]>
    <Arrow [2013-05-05T13:30:00+00:00]>
    <Arrow [2013-05-05T14:30:00+00:00]>
    <Arrow [2013-05-05T15:30:00+00:00]>
    <Arrow [2013-05-05T16:30:00+00:00]>
    """

if __name__ == '__main__':
    construct_demo()
    timezone_convert_demo()
    parse_format_demo()
    replace_shift_demo()
    properties_demo()
    replace_shift_demo()
    range_span_demo()