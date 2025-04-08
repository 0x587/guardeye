import json
import datetime
import streamlit as st
import pandas as pd

import grpc
from pb import link_pb2_grpc, link_pb2

channel = grpc.insecure_channel("10.0.4.112:8080")
cli = link_pb2_grpc.LinkStub(channel)


def main():
    st.title('Link')

    list_rsp: link_pb2.AgentListRsp = cli.AgentList(link_pb2.Empty())
    target_cid = st.selectbox(
        "选择Agent",
        list_rsp.agents.keys(),
        index=None,
    )
    st.write(f"Agent ID: {target_cid}")
    listen = cli.AgentListen(link_pb2.AgentListenReq(cid=target_cid))
    i = 0
    data = [{
        'topic': '',
        'json_data': {},
        'is_upstream': False,
        'datetime': datetime.datetime.now()
    }]
    df = st.dataframe(
        pd.DataFrame(
            data,
            columns=['topic', 'json_data', 'is_upstream', 'datetime'],
        ),
        use_container_width=True,
        column_config={
            "json_data": st.column_config.JsonColumn(
                "数据",
                width="large",
            ),
            "is_upstream": '是否上行',
            'datetime': '时间',
        },
        hide_index=True,
    )
    for a in listen:
        a: link_pb2.AgentListenRsp
        new_row = pd.DataFrame([{
            'topic': a.topic,
            'json_data': json.loads(a.json_data),
            'is_upstream': a.is_upstream,
            'datetime': datetime.datetime.now(),
        }], columns=['topic', 'json_data', 'is_upstream', 'datetime'])
        df.add_rows(new_row)
        i += 1


main()
