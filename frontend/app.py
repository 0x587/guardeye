import streamlit as st

st.set_page_config(layout='wide')

overview_page = st.Page('overview.py', title='Overview', url_path='/')
discover_page = st.Page('discover.py', title='Discover', url_path='/discover')
task_page = st.Page('task.py', title='Task', url_path='/task')

pg = st.navigation([overview_page, discover_page, task_page])
pg.run()
