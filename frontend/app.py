import streamlit as st

st.set_page_config(layout='wide')

overview_page = st.Page('overview.py', title='Overview')
discover_page = st.Page('discover.py', title='Discover')

pg = st.navigation([overview_page, discover_page])
pg.run()
