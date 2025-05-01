import ArrowBackIcon from '@mui/icons-material/ArrowBack'
import {
  Box,
  Button,
  Chip,
  styled,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography
} from '@mui/material'
import { green } from '@mui/material/colors'
import Image from 'next/image'
import { useRouter } from 'next/router'
import { MatchResult } from '../../components/MatchResult'
import { PageLayout } from '../../components/PageLayout'
import { Section } from '../../components/Section'
import { useHttp } from '../../hooks/useHttp'
import { fetcherStats } from '../../utils/http'
import { Match } from '../../utils/models'

function formatAction(playerName: string, action: string) {
  switch (action) {
    case 'goal':
      return `${playerName} fez um gol`
    case 'assist':
      return `${playerName} deu uma assistência`
    case 'yellow card':
      return `${playerName} levou um cartão amarelo`
    case 'red card':
      return `${playerName} levou um cartão vermelho`
    default:
      return `${playerName} fez alguma coisa`
  }
}

const StyledTableRow = styled(TableRow)(({ theme }) => ({
  '&>td': {
    border: 0
  },
  '&:nth-of-type(odd)': {
    backgroundColor: theme.palette.divider,
    border: 0
  }
}))

const StyledTableHead = styled(TableHead)({
  th: {
    border: 0
  }
})

const HeadCellContent = styled(Box)({
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center'
})

const StyledTableCell = styled(TableCell)(({ theme }) => ({
  textAlign: 'center'
}))

type HeadImageProps = {
  src: string
  alt: string
}

const HeadImage = (props: HeadImageProps) => (
  <Image src={props.src} alt={props.alt} width={32} height={32} />
)

const MatchPage = () => {
  const router = useRouter()
  const { id: matchId } = router.query
  const { data: match } = useHttp<Match>(
    matchId ? `/matches/${matchId}` : null,
    fetcherStats,
    { refreshInterval: 5000 }
  )

  return (
    <PageLayout>
      {match && (
        <Box
          sx={{
            display: 'flex',
            alignItems: 'center',
            flexDirection: 'column',
            gap: theme => theme.spacing(3)
          }}
        >
          <MatchResult match={match} />
          <Section
            sx={{
              marginTop: '-30px',
              zIndex: -10,
              width: 750,
              position: 'relative'
            }}
          >
            <Table sx={{ minWidth: 650 }} aria-label="simple table">
              <StyledTableHead>
                <TableRow>
                  <TableCell>
                    <HeadCellContent>
                      <HeadImage src="/img/time.svg" alt="" />
                      Tempo de jogo
                    </HeadCellContent>
                  </TableCell>
                  <TableCell>
                    <HeadCellContent>
                      <HeadImage src="/img/player.svg" alt="" /> Jogador
                    </HeadCellContent>
                  </TableCell>
                  <TableCell>
                    <HeadCellContent>
                      <HeadImage src="/img/score.svg" alt="" /> Pontuação
                    </HeadCellContent>
                  </TableCell>
                </TableRow>
              </StyledTableHead>
              <TableBody>
                {match.actions?.map((action, key) => (
                  <StyledTableRow key={key}>
                    <StyledTableCell>{action.minutes}&#39;</StyledTableCell>
                    <StyledTableCell>
                      {formatAction(action.player_name, action.action)}
                    </StyledTableCell>
                    <StyledTableCell
                      sx={{
                        color: theme =>
                          action.score > 0
                            ? green[500]
                            : theme.palette.primary.main
                      }}
                    >
                      <Typography>{action.score} pts</Typography>
                    </StyledTableCell>
                  </StyledTableRow>
                ))}
              </TableBody>
            </Table>
            <Chip
              label={
                <Box>
                  <Typography component="span">Total do jogo: </Typography>
                  <Typography
                    component="span"
                    sx={{
                      fontWeight: 'bold',
                      color: theme =>
                        1 > 0 ? green[500] : theme.palette.primary.main
                    }}
                  >
                    -- pts
                  </Typography>
                </Box>
              }
              sx={{
                bottom: -15,
                position: 'absolute',
                right: 15,
                backgroundColor: theme => theme.palette.background.default
              }}
            />
          </Section>
          <Button
            variant="contained"
            size="large"
            startIcon={<ArrowBackIcon />}
            onClick={() => router.back()}
          >
            Voltar
          </Button>
        </Box>
      )}
    </PageLayout>
  )
}

export default MatchPage
