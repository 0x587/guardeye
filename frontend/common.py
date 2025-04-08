import requests
import streamlit as st
import pandas as pd


def fetch_query(q, container=st, show_debug=False, title=None):
    if title: container.title(title)
    print('fCtx')
    res = requests.post(
        url='http://api.guardeye.shawnsiu.site:5080/api/v1/es/query',
        json={
            "query": q,
            "traceError": False
        }
    )
    print(res)
    if res.status_code != 200:
        container.error(res.content)
    rsp = res.json()
    if rsp['queryErrors'] is not None and len(rsp['queryErrors']) > 0:
        for e in rsp['queryErrors']:
            container.error(e)
    if show_debug:
        container.divider()
        profile = rsp['profile']
        profile['hitRate'] = f"{profile['resultCount'] / profile['fetchCount'] * 100:.2f}%"
        for c, (k, v) in zip(container.columns(len(profile)), profile.items()):
            c.metric(label=str(k), value=v)
    data = {i['timestamp']: i['value'] for i in rsp["data"]}
    df = pd.DataFrame.from_dict(data, orient="index", columns=rsp['columnNames'])
    df = df.astype(float)
    df.index = pd.to_datetime(df.index.astype(int), unit='s', utc=True).map(lambda x: x.tz_convert('Asia/Shanghai'))
    l, t = container.tabs(['line', 'table'])
    l.line_chart(df, x_label='time')
    t.dataframe(df)
