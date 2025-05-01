import { Avatar, ButtonProps, Chip } from '@mui/material'
import { AppBar, Box, Button, Toolbar } from '@mui/material'
import Image from 'next/image'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useHttp } from '../hooks/useHttp'
import { fetcherStats } from '../utils/http'
import { MY_TEAM_ID } from '../utils/models'

type NavBarItemProps = ButtonProps & {
  active: boolean
}

export const NavBarItem = (props: NavBarItemProps) => {
  const { children, active, ...rest } = props
  return (
    <Button
      LinkComponent={Link}
      sx={{
        color: 'white',
        display: 'inline-block',
        textAlign: 'center',
        '&::after': theme => ({
          content: '""',
          borderBottom: active
            ? `4px solid ${theme.palette.primary.main}`
            : `4px solid transparent`,
          width: '100%',
          display: 'block'
        })
      }}
      {...rest}
    >
      {children}
    </Button>
  )
}

export const NavBar = () => {
  const router = useRouter()

  const { data } = useHttp(`/my-teams/${MY_TEAM_ID}/balance`, fetcherStats, {
    refreshInterval: 5000
  })

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Image
            width={315}
            height={58}
            src="/img/logo.png"
            alt="Logo do Cartola FC"
            priority
          />

          <Box sx={{ flexGrow: 1, ml: theme => theme.spacing(4) }}>
            <NavBarItem active={router.pathname === '/'} href="/">
              Home
            </NavBarItem>
            <NavBarItem active={router.pathname === '/players'} href="/players">
              Escalação
            </NavBarItem>
            <NavBarItem
              active={['/matches', '/matches/[id]'].includes(router.pathname)}
              href="/matches"
            >
              Jogos
            </NavBarItem>
          </Box>

          <Chip
            label={data?.balance ?? 0}
            avatar={<Avatar>C$</Avatar>}
            color="secondary"
          />
        </Toolbar>
      </AppBar>
    </Box>
  )
}
