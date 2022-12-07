from django.contrib import admin

from .models import Action, Match, MyTeam, Player, Team

# Register your models here.
admin.site.register(Player)
admin.site.register(Team)
admin.site.register(MyTeam)
admin.site.register(Match)
admin.site.register(Action)
