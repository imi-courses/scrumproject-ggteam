import { FC, useEffect } from "react";
import {
  Page,
  Text,
  View,
  Document,
  StyleSheet,
  PDFViewer,
  Font,
} from "@react-pdf/renderer";
import { DialogContent } from "@/shared/ui/dialog";

Font.register({
  family: "Roboto",
  src: "https://fonts.gstatic.com/s/roboto/v20/KFOmCnqEu92Fr1Me5Q.ttf",
});

// Create styles
const styles = StyleSheet.create({
  page: {
    marginTop: 80,
    flexDirection: "column",
    backgroundColor: "#FFF",
    fontFamily: "Roboto",
    gap: 300,
  },
  section: {
    margin: 10,
  },
  writer: {
    margin: 10,
    padding: 10,
    width: "100%",
  },
  sectionData: {
    marginLeft: 40,
    padding: 4,
    flexDirection: "row",
    fontSize: "14px",
    width: "100%",
  },
  sectionDataReal: {
    marginLeft: 40,
    padding: 4,
    flexDirection: "row",
    fontSize: "14px",
    width: "100%",
    marginTop: 24,
  },
  viewer: {
    position: "absolute",
    top: "50%",
    left: "50%",
    marginRight: "-50%",
    transform: "translate(-50%, -50%)",
    width: window.innerWidth / 2,
    height: window.innerHeight / 1.2,
  },
  title: {
    textAlign: "center",
    marginBottom: 20,
  },
  textData: {
    textDecoration: "underline",
  },
  textDataFill: {
    display: "flex",
    flexDirection: "row",
    width: "70%",
    borderBottom: "1 solid black",
  },
  textWriter: {
    marginLeft: 20,
    marginBottom: 12,
    fontSize: 14,
  },
  textDataFillTitle: {
    width: 50,
  },
  textSign: {
    borderBottom: "1 solid black",
    width: 128,
    textAlign: "right",
    marginRight: 84,
  },
  sign: {
    fontSize: 14,
  },
  textDate: {
    borderBottom: "1 solid black",
    width: 80,
    textAlign: "right",
  },
  date: {
    fontSize: 14,
  },
  viewSign: {
    width: "100%",
    alignItems: "flex-end",
    justifyContent: "flex-end",
    flexDirection: "row",
    gap: 40,
  },
  viewMy: {
    flexDirection: "row",
    gap: 8,
  },
});

interface CreateRealEstatePDFProps {
  fio: string;
  email: string;
  phone: string;
}

// Create Document Component
const CreateRealEstatePDF: FC<CreateRealEstatePDFProps> = (props) => {
  const { fio, email, phone } = props;
  useEffect(() => {
    async () => {
      await Font.load({ fontFamily: "Roboto" });
    };
  }, []);
  return (
    <DialogContent
      datatype="pdf"
      className="bg-transparent border-0 shadow-transparent"
    >
      <PDFViewer style={styles.viewer}>
        <Document>
          <Page size="A4" style={styles.page}>
            <View>
              <View style={styles.section}>
                <Text style={styles.title}>
                  Заявление о регистрации недвижимости
                </Text>
              </View>
              <View style={styles.writer}>
                <Text style={styles.textWriter}>Заявитель:</Text>
                <View style={styles.sectionData}>
                  <Text>ФИО: </Text>
                  <Text style={styles.textData}>{fio}</Text>
                </View>
                <View style={styles.sectionData}>
                  <Text>Номер телефона: </Text>
                  <Text style={styles.textData}>{phone}</Text>
                </View>
                <View style={styles.sectionData}>
                  <Text>Почта: </Text>
                  <Text style={styles.textData}>{email}</Text>
                </View>
              </View>
              <View style={styles.writer}>
                <Text style={styles.textWriter}>Недвижимость:</Text>
                <View style={styles.sectionDataReal}>
                  <Text style={styles.textDataFillTitle}>Адрес: </Text>
                  <Text style={styles.textDataFill}></Text>
                </View>
                <View style={styles.sectionDataReal}>
                  <Text style={styles.textDataFillTitle}>Тип: </Text>
                  <Text style={styles.textDataFill}></Text>
                </View>
              </View>
            </View>
            <View style={styles.viewSign}>
              <View style={styles.viewMy}>
                <Text style={styles.date}>Дата: </Text>
                <Text style={styles.textDate}></Text>
              </View>
              <View style={styles.viewMy}>
                <Text style={styles.sign}>Подпись:</Text>
                <Text style={styles.textSign}></Text>
              </View>
            </View>
          </Page>
        </Document>
      </PDFViewer>
    </DialogContent>
  );
};
export default CreateRealEstatePDF;
