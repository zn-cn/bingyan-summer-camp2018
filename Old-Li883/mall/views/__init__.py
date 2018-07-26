from flask import Flask

app = Flask(__name__)
import views.client_views
import views.guys_views
import views.merchant_views
import views.administration_views