import Layout from './Layout'
import { BrowserRouter } from 'react-router-dom'
import { test } from 'vitest'

test('Renders content', async () => {
  return (
    <BrowserRouter>
      <Layout>
        <div></div>
      </Layout>
    </BrowserRouter>
  )
})
