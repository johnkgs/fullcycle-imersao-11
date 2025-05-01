import GroupsIcon from '@mui/icons-material/Groups'
import { Button, Divider, Grid, styled } from '@mui/material'
import Link from 'next/link'
import { Label } from '../components/Label'
import { PageLayout } from '../components/PageLayout'
import { Section } from '../components/Section'
import { TeamLogo } from '../components/TeamLogo'
import { useHttp } from '../hooks/useHttp'
import { fetcherStats } from '../utils/http'
import { MY_TEAM_ID } from '../utils/models'

const BudgetContainer = styled(Section)(({ theme }) => ({
  width: '800px',
  height: '300px',
  marginTop: theme.spacing(8),
  display: 'flex',
  alignItems: 'center'
}))

const HomePage = () => {
  const { data } = useHttp(`/my-teams/${MY_TEAM_ID}/balance`, fetcherStats, {
    refreshInterval: 5000
  })

  return (
    <PageLayout>
      <Grid
        container
        sx={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          gap: theme => theme.spacing(3)
        }}
      >
        <Grid item>
          <TeamLogo
            sx={{ position: 'absolute', left: 0, right: 0, m: 'auto' }}
          />
          <BudgetContainer>
            <Grid container>
              <Grid
                item
                xs={5}
                sx={{
                  display: 'flex',
                  flexDirection: 'column',
                  alignItems: 'center'
                }}
              >
                <Label>Última pontuação</Label>
                <Label>-</Label>
              </Grid>
              <Grid
                item
                xs={2}
                sx={{ display: 'flex', justifyContent: 'center' }}
              >
                <Divider orientation="vertical" sx={{ height: 'auto' }} />
              </Grid>
              <Grid
                item
                xs={5}
                sx={{
                  display: 'flex',
                  flexDirection: 'column',
                  alignItems: 'center'
                }}
              >
                <Label>Patrimônio</Label>
                <Label>{data?.balance ?? 0}</Label>
              </Grid>
            </Grid>
          </BudgetContainer>
        </Grid>
        <Grid item>
          <Button
            component={Link}
            href="/players"
            variant="contained"
            startIcon={<GroupsIcon />}
          >
            Escalar Jogadores
          </Button>
        </Grid>
      </Grid>
    </PageLayout>
  )
}

export default HomePage
