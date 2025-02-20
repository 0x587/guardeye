import streamlit as st
from streamlit_ace import st_ace
from common import fetch_query

q1 = \
    """select 
    json($msg).'data'.'angleA' with avg as angleA,
    json($msg).'data'.'angleB' as angleB,
    json($msg).'data'.'angleC' as angleC
window 1m
from node * provider mqtt
in -1d;"""



st.title('Discover')
query = st_ace(language='sql', value=q1, theme='tomorrow_night_eighties')

if query:
    fetch_query(query)
