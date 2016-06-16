package asciiFolding

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestTextFolding(t *testing.T) {
    assert := assert.New(t)
    
    assert.Equal("Some", Fold("Söme"))
    assert.Equal("Gum", Fold("Güm"))
    assert.Equal("Cunku", Fold("Çünkü"))
    assert.Equal("alemi huzun icundur", Fold("âlêmî hüzün içûndur"))
    assert.Equal("Passe compose", Fold("Passé composé"))
    assert.Equal("En reponse a votre annonce dans le magazine FUSAC le 20 juillet 2010, je vous ecris pour poser ma candidature pour le poste de enseignante remplacant d'anglais dans votre ecole.  Puisque je m'interesse beaucoup a l'enseignement des langues, ce poste correspond parfaitement a mes objectifs professionnels.", Fold("En réponse à votre annonce dans le magazine FUSAC le 20 juillet 2010, je vous écris pour poser ma candidature pour le poste de enseignante remplaçant d’anglais dans votre école.  Puisque je m’intéresse beaucoup à l’enseignement des langues, ce poste correspond parfaitement à mes objectifs professionnels."))
    assert.Equal("Assistante d'anglais dans des ecoles primaires en France pendant deux ans (Academie de Paris en 2007-2008 et Academie de Creteil en 2009-2010), j'ai egalement cree un cours intensif de langue francaise pour des enfants ages de 4 a 9 ans dans le Massachusetts en 2008.  Puisque j'emploie beaucoup de methodologies diverses dans l'organisation de mes cours, ces eleves ont appris par coeur des poemes tels que \" Sonnet 116 \" de Shakespeare et \" Dreams \" de Langston Hughes, et ils ont chante \" Imagine \" des Beatles lors de leur spectacle de fin d'annee.  Dans les petites classes, l'usage de la chanson et du jeu ont ete d'une importance capitale pour engager les eleves et leur donner envie d'apprendre.", Fold("Assistante d’anglais dans des écoles primaires en France pendant deux ans (Académie de Paris en 2007-2008 et Académie de Créteil en 2009-2010), j’ai également créé un cours intensif de langue française pour des enfants âgés de 4 à 9 ans dans le Massachusetts en 2008.  Puisque j’emploie beaucoup de méthodologies diverses dans l’organisation de mes cours, ces élèves ont appris par cœur des poèmes tels que « Sonnet 116 » de Shakespeare et « Dreams » de Langston Hughes, et ils ont chanté « Imagine » des Beatles lors de leur spectacle de fin d’année.  Dans les petites classes, l’usage de la chanson et du jeu ont été d’une importance capitale pour engager les élèves et leur donner envie d’apprendre."))
}