import type { NextPage } from 'next'
import React, { useMemo, useState } from 'react';
import Head from 'next/head'
import homeStyle from '../styles/home.module.css'
import articleStyle from '../styles/article.module.css'
const Home: NextPage = () => {

  return (
    <>
      <Head>
        <title>ZZH Blog</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div  className={articleStyle.articleList} data-layout="article">
          <div className={articleStyle.articleItem} data-layout="article-item">1231</div>
          <div className={articleStyle.articleItem} data-layout="article-item">1231</div>
          <div className={articleStyle.articleItem} data-layout="article-item">1231</div>
          <div className={articleStyle.articleItem} data-layout="article-item">1231</div>
          <div className={articleStyle.articleItem} data-layout="article-item">1231</div>
          <div className={articleStyle.articleItem} data-layout="article-item">1231</div>
      </div>
    </>

  )
}

export default Home
