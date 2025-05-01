import { Box } from '@mui/material'
import { useRouter } from 'next/router'
import { MatchResult } from '../../components/MatchResult'
import { PageLayout } from '../../components/PageLayout'
import { useHttp } from '../../hooks/useHttp'
import { fetcherStats } from '../../utils/http'
import { Match } from '../../utils/models'

const MatchesPage = () => {
  const { data } = useHttp<Match[]>('/matches', fetcherStats, {
    refreshInterval: 5000
  })
  const router = useRouter()
  return (
    <PageLayout>
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
          flexDirection: 'column',
          gap: theme => theme.spacing(3)
        }}
      >
        {data &&
          data.map((match, key) => (
            <Box
              key={key}
              sx={{ cursor: 'pointer' }}
              onClick={() => router.push(`/matches/${match.id}`)}
            >
              <MatchResult match={match} />
            </Box>
          ))}
      </Box>
    </PageLayout>
  )
}

export default MatchesPage
