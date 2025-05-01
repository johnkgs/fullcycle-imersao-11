import { Container } from '@mui/material'
import type { PropsWithChildren } from 'react'

export const PageLayout = (props: PropsWithChildren) => {
  const { children } = props
  return (
    <Container
      sx={{
        pt: theme => theme.spacing(3)
      }}
    >
      {children}
    </Container>
  )
}
