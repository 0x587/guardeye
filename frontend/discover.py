import time

import requests
import streamlit as st
from streamlit_ace import st_ace
from streamlit_autorefresh import st_autorefresh
from common import fetch_query

q1 = \
    """select 
    json($msg).'data'.'angleA' with avg as angleA,
    json($msg).'data'.'angleB' as angleB,
    json($msg).'data'.'angleC' as angleC
window 1m
from node * provider mqtt
in -1h;"""

st.title('Discover')
query = st_ace(language='sql', value=q1, theme='tomorrow_night_eighties')

if st.button('fetch'):
    fetch_query(query)

if st.button('export'):
    res = requests.post(
        url='http://api.guardeye.shawnsiu.site:5080/api/v1/es/export',
        json={
            "query": query,
            "traceError": False
        }
    )
    task_id = res.json()['taskId']
    st.write(fCtx"TaskID: {task_id}")
    st.session_state['task_id'] = task_id
    st.switch_page('task.py')

    # count = st_autorefresh(interval=1000, limit=None, key="fizzbuzzcounter")
    # if count >= 0:
    # if st.button('check'):
    #     res = requests.post(
    #         url='http://localhost:8888/api/v1/es/taskstatus',
    #         json={
    #             "taskId": task_id,
    #         }
    #     )
    #     st.write(res.json())
