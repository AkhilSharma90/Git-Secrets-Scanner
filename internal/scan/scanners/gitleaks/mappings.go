package gitleaks

import "github.com/AkhilSharma90/Git-Secrets-Scanner/internal/report/secret"

var GitleaksSecretKindMapping = map[string]secret.SecretKind{
	"1password-service-account-token":    secret.SecretKind1Password,
	"adafruit-api-key":                   secret.SecretKindAdafruitIO,
	"adobe-client-id":                    secret.SecretKindAdobeIO,
	"adobe-client-secret":                secret.SecretKindAdobeIO,
	"age-secret-key":                     secret.SecretKindAge,
	"airtable-api-key":                   secret.SecretKindAirtable,
	"algolia-api-key":                    secret.SecretKindAlgolia,
	"alibaba-access-key-id":              secret.SecretKindAlibaba,
	"alibaba-secret-key":                 secret.SecretKindAlibaba,
	"asana-client-id":                    secret.SecretKindAsana,
	"asana-client-secret":                secret.SecretKindAsana,
	"atlassian-api-token":                secret.SecretKindAtlassian,
	"authress-service-client-access-key": secret.SecretKindAuthress,
	"aws-access-token":                   secret.SecretKindAWS,
	"azure-ad-client-secret":             secret.SecretKindAzure,
	"beamer-api-token":                   secret.SecretKindBeamer,
	"bitbucket-client-id":                secret.SecretKindBitBucket,
	"bitbucket-client-secret":            secret.SecretKindBitBucket,
	"bittrex-access-key":                 secret.SecretKindBittrex,
	"bittrex-secret-key":                 secret.SecretKindBittrex,
	"clojars-api-token":                  secret.SecretKindClojars,
	"cloudflare-api-key":                 secret.SecretKindCloudflare,
	"cloudflare-global-api-key":          secret.SecretKindCloudflare,
	"cloudflare-origin-ca-key":           secret.SecretKindCloudflare,
	"codecov-access-token":               secret.SecretKindCodecov,
	"cohere-api-token":                   secret.SecretKindCohere,
	"coinbase-access-token":              secret.SecretKindCoinbase,
	"confluent-access-token":             secret.SecretKindConfluent,
	"confluent-secret-key":               secret.SecretKindConfluent,
	"contentful-delivery-api-token":      secret.SecretKindConfluent,
	"databricks-api-token":               secret.SecretKindDatabricks,
	"datadog-access-token":               secret.SecretKindDatadog,
	"defined-networking-api-token":       secret.SecretKindDefinedNetworking,
	"digitalocean-access-token":          secret.SecretKindDigitalOcean,
	"digitalocean-pat":                   secret.SecretKindDigitalOcean,
	"digitalocean-refresh-token":         secret.SecretKindDigitalOcean,
	"discord-api-token":                  secret.SecretKindDiscord,
	"discord-client-id":                  secret.SecretKindDiscord,
	"discord-client-secret":              secret.SecretKindDiscord,
	"doppler-api-token":                  secret.SecretKindDoppler,
	"droneci-access-token":               secret.SecretKindDroneCI,
	"dropbox-api-token":                  secret.SecretKindDropbox,
	"dropbox-long-lived-api-token":       secret.SecretKindDropbox,
	"dropbox-short-lived-api-token":      secret.SecretKindDropbox,
	"duffel-api-token":                   secret.SecretKindDuffel,
	"dynatrace-api-token":                secret.SecretKindDynatrace,
	"easypost-api-token":                 secret.SecretKindGeneric,
	"easypost-test-api-token":            secret.SecretKindGeneric,
	"etsy-access-token":                  secret.SecretKindEtsy,
	"facebook-access-token":              secret.SecretKindFacebook,
	"facebook-page-access-token":         secret.SecretKindFacebook,
	"facebook-secret":                    secret.SecretKindFacebook,
	"fastly-api-token":                   secret.SecretKindFastly,
	"finicity-api-token":                 secret.SecretKindFinicity,
	"finicity-client-secret":             secret.SecretKindFinicity,
	"finnhub-access-token":               secret.SecretKindFinnhub,
	"flickr-access-token":                secret.SecretKindFlickr,
	"flutterwave-encryption-key":         secret.SecretKindFlutterwave,
	"flutterwave-public-key":             secret.SecretKindFlutterwave,
	"flutterwave-secret-key":             secret.SecretKindFlutterwave,
	"flyio-access-token":                 secret.SecretKindFlyIO,
	"frameio-api-token":                  secret.SecretKindFrameIO,
	"freshbooks-access-token":            secret.SecretKindFreshbooks,
	"gcp-api-key":                        secret.SecretKindGCP,
	"generic-api-key":                    secret.SecretKindGeneric,
	"github-app-token":                   secret.SecretKindGithub,
	"github-fine-grained-pat":            secret.SecretKindGithub,
	"github-oauth":                       secret.SecretKindGithub,
	"github-pat":                         secret.SecretKindGithub,
	"github-refresh-token":               secret.SecretKindGithub,
	"gitlab-pat":                         secret.SecretKindGitlab,
	"gitlab-ptt":                         secret.SecretKindGitlab,
	"gitlab-rrt":                         secret.SecretKindGitlab,
	"gitter-access-token":                secret.SecretKindGitter,
	"gocardless-api-token":               secret.SecretKindGoCardless,
	"grafana-api-key":                    secret.SecretKindGrafana,
	"grafana-cloud-api-token":            secret.SecretKindGrafana,
	"grafana-service-account-token":      secret.SecretKindGrafana,
	"harness-api-key":                    secret.SecretKindHarness,
	"hashicorp-tf-api-token":             secret.SecretKindTerraformCloud,
	"heroku-api-key":                     secret.SecretKindHeroku,
	"hubspot-api-key":                    secret.SecretKindHubSpot,
	"huggingface-access-token":           secret.SecretKindHuggingFace,
	"huggingface-organization-api-token": secret.SecretKindHuggingFace,
	"infracost-api-token":                secret.SecretKindInfracost,
	"intercom-api-key":                   secret.SecretKindIntercom,
	"intra42-client-secret":              secret.SecretKindIntra42,
	"jfrog-api-key":                      secret.SecretKindJFrog,
	"jfrog-identity-token":               secret.SecretKindJFrog,
	"jwt-base64":                         secret.SecretKindJWT,
	"jwt":                                secret.SecretKindJWT,
	"kraken-access-token":                secret.SecretKindKraken,
	"kucoin-access-token":                secret.SecretKindKuCoin,
	"kucoin-secret-key":                  secret.SecretKindKuCoin,
	"launchdarkly-access-token":          secret.SecretKindLaunchDarkly,
	"linear-api-key":                     secret.SecretKindLinear,
	"linear-client-secret":               secret.SecretKindLinear,
	"linkedin-client-id":                 secret.SecretKindLinkedIn,
	"linkedin-client-secret":             secret.SecretKindLinkedIn,
	"lob-api-key":                        secret.SecretKindLob,
	"lob-pub-api-key":                    secret.SecretKindLob,
	"mailchimp-api-key":                  secret.SecretKindMailchimp,
	"mailgun-private-api-token":          secret.SecretKindMailgun,
	"mailgun-pub-key":                    secret.SecretKindMailgun,
	"mailgun-signing-key":                secret.SecretKindMailgun,
	"mapbox-api-token":                   secret.SecretKindMapBox,
	"mattermost-access-token":            secret.SecretKindMattermost,
	"messagebird-api-token":              secret.SecretKindMessageBird,
	"messagebird-client-id":              secret.SecretKindMessageBird,
	"microsoft-teams-webhook":            secret.SecretKindMicrosoftTeams,
	"netlify-access-token":               secret.SecretKindNetlify,
	"new-relic-browser-api-token":        secret.SecretKindNewRelic,
	"new-relic-insert-key":               secret.SecretKindNewRelic,
	"new-relic-user-api-id":              secret.SecretKindNewRelic,
	"new-relic-user-api-key":             secret.SecretKindNewRelic,
	"npm-access-token":                   secret.SecretKindNpm,
	"nuget-config-password":              secret.SecretKindNuGet,
	"nytimes-access-token":               secret.SecretKindNytimes,
	"okta-access-token":                  secret.SecretKindOkta,
	"openai-api-key":                     secret.SecretKindOpenAI,
	"openshift-user-token":               secret.SecretKindOpenShift,
	"plaid-api-token":                    secret.SecretKindPlaid,
	"plaid-client-id":                    secret.SecretKindPlaid,
	"plaid-secret-key":                   secret.SecretKindPlaid,
	"planetscale-api-token":              secret.SecretKindPlanetScale,
	"planetscale-oauth-token":            secret.SecretKindPlanetScale,
	"planetscale-password":               secret.SecretKindPlanetScale,
	"postman-api-token":                  secret.SecretKindPostman,
	"prefect-api-token":                  secret.SecretKindPrefect,
	"private-key":                        secret.SecretKindPrivateKey,
	"privateai-api-token":                secret.SecretKindPrivateAI,
	"pulumi-api-token":                   secret.SecretKindPulumi,
	"pypi-upload-token":                  secret.SecretKindPyPI,
	"rapidapi-access-token":              secret.SecretKindRapidApi,
	"readme-api-token":                   secret.SecretKindReadMe,
	"rubygems-api-token":                 secret.SecretKindRubyGems,
	"scalingo-api-token":                 secret.SecretKindScalingo,
	"sendbird-access-id":                 secret.SecretKindSendbird,
	"sendbird-access-token":              secret.SecretKindSendbird,
	"sendgrid-api-token":                 secret.SecretKindSendGrid,
	"sendinblue-api-token":               secret.SecretKindSendinBlue,
	"sentry-access-token":                secret.SecretKindSentry,
	"shippo-api-token":                   secret.SecretKindShippo,
	"shopify-access-token":               secret.SecretKindShopify,
	"shopify-custom-access-token":        secret.SecretKindShopify,
	"shopify-private-app-access-token":   secret.SecretKindShopify,
	"shopify-shared-secret":              secret.SecretKindShopify,
	"sidekiq-secret":                     secret.SecretKindSidekiq,
	"sidekiq-sensitive-url":              secret.SecretKindSidekiq,
	"slack-app-token":                    secret.SecretKindSlack,
	"slack-bot-token":                    secret.SecretKindSlack,
	"slack-config-access-token":          secret.SecretKindSlack,
	"slack-config-refresh-token":         secret.SecretKindSlack,
	"slack-legacy-bot-token":             secret.SecretKindSlack,
	"slack-legacy-token":                 secret.SecretKindSlack,
	"slack-legacy-workspace-token":       secret.SecretKindSlack,
	"slack-user-token":                   secret.SecretKindSlack,
	"slack-webhook-url":                  secret.SecretKindSlack,
	"snyk-api-token":                     secret.SecretKindSnyk,
	"square-access-token":                secret.SecretKindSquare,
	"squarespace-access-token":           secret.SecretKindSquarespace,
	"stripe-access-token":                secret.SecretKindStripe,
	"sumologic-access-id":                secret.SecretKindSumoLogic,
	"sumologic-access-token":             secret.SecretKindSumoLogic,
	"telegram-bot-api-token":             secret.SecretKindTelegram,
	"travisci-access-token":              secret.SecretKindTravisCI,
	"twilio-api-key":                     secret.SecretKindTwilio,
	"twitch-api-token":                   secret.SecretKindTwitch,
	"twitter-access-secret":              secret.SecretKindTwitter,
	"twitter-access-token":               secret.SecretKindTwitter,
	"twitter-api-key":                    secret.SecretKindTwitter,
	"twitter-api-secret":                 secret.SecretKindTwitter,
	"twitter-bearer-token":               secret.SecretKindTwitter,
	"typeform-api-token":                 secret.SecretKindTypeform,
	"vault-batch-token":                  secret.SecretKindVault,
	"vault-service-token":                secret.SecretKindVault,
	"yandex-access-token":                secret.SecretKindYandex,
	"yandex-api-key":                     secret.SecretKindYandex,
	"yandex-aws-access-token":            secret.SecretKindYandex,
	"zendesk-secret-key":                 secret.SecretKindZendesk,

	// ignore generic secrets found by some custom gitleaks detectors
	"curl-auth-header":       secret.SecretKindGeneric,
	"curl-auth-user":         secret.SecretKindGeneric,
	"hashicorp-tf-password":  secret.SecretKindGeneric,
	"kubernetes-secret-yaml": secret.SecretKindGeneric,
}