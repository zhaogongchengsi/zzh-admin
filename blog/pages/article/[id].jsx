import { useRouter } from 'next/router'
import Editor from 'md-editor-rt'
import 'md-editor-rt/lib/style.css';
import { get, confirm } from '../../api/index'

const Article = (props) => {
    const { article_context } = props.article

  return <>
    <div className="article-prev">
        <Editor previewOnly={true} editorClass="ed-my" previewTheme="vuepress" theme="dark" modelValue={article_context} />
    </div>
  </>
}


export async function getServerSideProps({params}) {
    const _res = await get(`article/get_article_text?id=${params.id}`)
    const res = confirm(_res)
    if (res) {
      return {
        props: {
          article:res
        }, // will be passed to the page component as props
      }
    }
    return {notFound:true}
  }

export default Article