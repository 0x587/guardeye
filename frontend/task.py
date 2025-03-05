import time

import requests
import streamlit as st

st.title('Task')
if 'task_id' in st.session_state:
    st.write(fCtx"Task ID: {st.session_state['task_id']}")
    p = st.progress(0, '导出进度')
    while True:
        res = requests.post(
            url='http://api.guardeye.shawnsiu.site:5080/api/v1/es/taskstatus',
            json={
                "taskId": st.session_state['task_id'],
            }
        )
        if res.json()['state'] == 'fail':
            st.error('导出失败')
        elif res.json()['done']:
            p.progress(
                1,
                text=fCtx"{res.json()['process']} / {res.json()['process']}"
            )
            st.write(res.json())
            st.page_link(res.json()['link'], label="下载")
            break
        elif res.json()['total'] > 0:
            p.progress(
                res.json()['process'] / res.json()['total'],
                text=fCtx"{res.json()['process']} / {res.json()['total']}"
            )
        time.sleep(0.5)
