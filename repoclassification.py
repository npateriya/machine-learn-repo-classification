from sklearn.pipeline import Pipeline
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.pipeline import Pipeline
from sklearn.feature_extraction.text import TfidfTransformer
from sklearn.naive_bayes import MultinomialNB
from sklearn.linear_model import SGDClassifier
import sklearn.datasets


repo_train = sklearn.datasets.load_files("./TrainReadmes")
#print repo_train.target_names #prints all the categories
#print (repo_train.target) #prints all the categories
#print len(repo_train.data)

repo_test = sklearn.datasets.load_files("./Readmes")
#print len(repo_test.data)

text_clf = Pipeline([('vect', CountVectorizer(stop_words='english')),
	 ('tfidf', TfidfTransformer()), 
	 #('clf', MultinomialNB()) ])
	 ('clf-svm', SGDClassifier(loss="log", penalty='l2',alpha=1e-3, random_state=42))])

text_clf = text_clf.fit(repo_train.data, repo_train.target)

predicted = text_clf.predict(repo_test.data)
prob = text_clf.predict_proba(repo_test.data)
classification = {}
for i in range(0, len(repo_test.data)-1):
	org, repo = repo_test.filenames[i].split("__")
	org = org.replace("./Readmes/unknown/", "")
	repo_name = "https://github.com/"+org+"/"+repo
	classification[repo_name]={"url":repo_name, "tag":repo_train.target_names[predicted[i]], "probability": prob[i][predicted[i]]}
	#print(repo_name +","+repo_train.target_names[predicted[i]] +"," + str(prob[i][predicted[i]]))

for k, v in sorted(classification.items(), key=lambda x: x[1]['probability'], reverse=True):
	    print k," , ",  v['probability'], " , ", v['tag']
