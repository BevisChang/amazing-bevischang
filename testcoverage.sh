basicCoverageRate=700
repositoryName="bevis-chang"
coverageResult=$(go tool cover -func=coverage.out)

# make sure this service template name wont replase by one shot init script
serviceTemplateRepoName="go"
serviceTemplateRepoName="${serviceTemplateRepoName}-amazing"

#  allow service template only have 20% test coverage
if [ $repositoryName = $serviceTemplateRepoName ] 
then 
  basicCoverageRate=200
fi

echo "$coverageResult\n-----------------------"

coverageRate=$(go tool cover -func=coverage.out | grep "(statements" | sed 's/[^0-9]*//g');

if [ $coverageRate -gt $basicCoverageRate ]
then
  echo "✅   test coverage rate is higher than 70%"
  exit 0 
else
  echo "⚠️   test coverage rate should higher than 70%"
  exit 1
fi