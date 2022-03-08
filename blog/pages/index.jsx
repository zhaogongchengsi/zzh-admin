import React, { useMemo, useState } from 'react'
import Head from 'next/head'
import { useRouter } from 'next/router'
import articleStyle from '../styles/article.module.css'
import { get, confirm } from '../api/index'
import { Pagination } from '@douyinfe/semi-ui'

const Home = (props) => {
  const { article_list, count, limit_offset } = props.articles
  const router = useRouter()

  const postRouter = (id) => {
    router.push(`article/${id}`)
  }

  return (
    <>
      <Head>
        <title>ZZH Blog</title>
        <link rel="icon" href="/zzh_logo.png" />
      </Head>
      <div className={articleStyle.articleList} data-layout="article">
        {article_list.map((item) => {
          return (
            <div
              className={`${articleStyle.articleItem} dark:bg-gray-600 dark:text-white`}
              data-layout="article-item"
              key={item.ID}
              onClick={() => {
                postRouter(item.ID)
              }}
            >
              {item.article_title}
            </div>
          )
        })}
      </div>
      <div className="pagination flex items-center justify-center p-2.5">
        <Pagination
          currentPage={limit_offset.offset}
          pageSize={limit_offset.limit}
          total={count}
        />
      </div>
    </>
  )
}

export async function getServerSideProps() {
  const _res = await get(`article/get_article_list?limit=${10}&offset=${0}`)
  const res = confirm(_res)
  if (res) {
    return {
      props: {
        articles: res,
      }, // will be passed to the page component as props
    }
  }
  return { notFound: true }
}

export default Home
