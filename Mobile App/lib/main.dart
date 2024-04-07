import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:flutter_colorpicker/flutter_colorpicker.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'VendIQ',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: SoundIntensityPage(),
      debugShowCheckedModeBanner: false,
    );
  }
}

class SoundIntensityPage extends StatefulWidget {
  @override
  _SoundIntensityPageState createState() => _SoundIntensityPageState();
}

class _SoundIntensityPageState extends State<SoundIntensityPage> {
  String ipAddress = '';
  Color selectedColor = Colors.blue;
  double soundIntensity = 0.0;
  bool isSendingAudio = false;

  void startSendingAudio() async {
    // Simulate sending audio data

    // Send POST request to your server
    final response = await http.post(
      Uri.parse('https://tunnel.singhropar.tech/update-message'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'message': 'Connected',
      }),
    );

    if (response.statusCode == 200) {
      print('Audio data sent successfully');
    } else {
      print('Failed to send audio data: ${response.reasonPhrase}');
    }

    @override
    Widget build(BuildContext context) {
      return Scaffold(
        appBar: AppBar(
          title: Text('VendIQ'),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              ElevatedButton(
                onPressed: isSendingAudio ? null : startSendingAudio,
                child: Text(isSendingAudio
                    ? 'Sending Audio...'
                    : 'Start Calling Server'),
              ),
            ],
          ),
        ),
      );
    }
  }
  
  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    throw UnimplementedError();
  }
}
