import streamlit as st
from common import fetch_query

q1 = \
    """select 
    json($msg).'data'.'angleA' with avg as angleA,
    json($msg).'data'.'angleB' with avg as angleB,
    json($msg).'data'.'angleC' with avg as angleC
window 1m
from node * provider mqtt
in -1d;"""

q2 = \
    """select 
    json($msg).'data'.'voltageA' with avg as voltageA,
    json($msg).'data'.'voltageB' with avg as voltageB,
    json($msg).'data'.'voltageC' with avg as voltageC
window 1m
from node * provider mqtt
in -1d;"""

st.title('Overview')
l, r = st.columns(2)

fetch_query(q1, l, title="Angle")
fetch_query(q2, r, title="Voltage")
