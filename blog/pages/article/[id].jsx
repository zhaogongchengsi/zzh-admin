import { useRouter } from 'next/router'
import Editor from 'md-editor-rt'
import 'md-editor-rt/lib/style.css'
import { get, confirm } from '../../api/index'
import { useState } from 'react'

const Article = (props) => {
  const { article_context } = props.article
  const [conList, setconlist] = useState([])

  // theme = 'dark' 黑夜模式

  const getCatalog = (list) => {
    setconlist(list)
  }

  return (
    <>
      <div className="article-prev">
        <Editor
          previewOnly={true}
          editorClass="ed-my rounded-lg border-4 border-blue-700 border-double"
          previewTheme="vuepress"
          modelValue={article_context}
          onGetCatalog={getCatalog}
        />
        <div className="content-list sticky rounded shadow hover:shadow-md dark:bg-gray-900">
          <span className="block text-xl font-semibold dark:text-white">
            目录
          </span>
          {conList.map((item, index) => {
            return (
              <a
                href={`#${item.text}`}
                className="my-1.5 block truncate text-sm hover:text-green-600 dark:text-white"
                key={index}
              >
                {item.text}
              </a>
            )
          })}
        </div>
      </div>
    </>
  )
}

export async function getServerSideProps({ params }) {
  const _res = await get(`article/get_article_text?id=${params.id}`)
  const res = confirm(_res)
  if (res) {
    return {
      props: {
        article: res,
      },
    }
  }
  return { notFound: true }
}

export default Article
